package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/GlobalCents-DocumentFactory/mcp-documentfactory/internal/model"
	"github.com/google/uuid"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/nats-io/nats.go"
)

const (
	natsServerName = "df-mcp-server"
	httpPort       = ":8080"
)

// Config holds the application configuration.
type Config struct {
	NatsURL    string
	LicenseKey string
	IsLicensed bool
}

// DownloadToken represents a secure, one-time-use token for downloading a file.
type DownloadToken struct {
	Bucket    string
	Key       string
	ExpiresAt time.Time
}

// TokenStore holds the active download tokens.
var (
	tokenStore = make(map[string]DownloadToken)
	tokenMutex = &sync.RWMutex{}
)

// getEnvOrDefault returns environment variable value or default if not set
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// isValidLicense validates the license key format and authenticity
func isValidLicense(licenseKey string) bool {
	// Basic validation - in production this would verify against a license server
	// For now, accept any non-empty license key that looks valid
	if len(licenseKey) < 8 {
		return false
	}

	// Check for basic format (could be enhanced with cryptographic validation)
	if strings.Contains(licenseKey, " ") || len(licenseKey) > 128 {
		return false
	}

	// In production, you would validate against a license server or use
	// cryptographic signatures to verify the license authenticity
	return true
}

func main() {
	// Start the HTTP server for downloads in a separate goroutine
	go startDownloadServer()

	s := server.NewMCPServer("DocumentFactory MCP Server", "1.0.0",
		server.WithToolCapabilities(true),
		server.WithLogging(),
	)

	cfg := &Config{
		NatsURL:    getEnvOrDefault("NATS_URL", "nats://nats.documentfactory.com:4222"),
		LicenseKey: os.Getenv("LICENCEKEY"),
	}
	cfg.IsLicensed = cfg.LicenseKey != "" && isValidLicense(cfg.LicenseKey)

	// Log mode for debugging
	if cfg.IsLicensed {
		log.Printf("Running in LICENSED mode")
	} else {
		log.Printf("Running in DEMO mode (watermarks will be applied)")
	}

	s.AddTool(
		mcp.NewTool("pipeline",
			mcp.WithDescription("The 'pipeline' tool allows you to define and execute a series of document processing steps like converting formats, adding watermarks, or merging documents. The pipeline is defined by a JSON object."),
			mcp.WithArray("files",
				mcp.Required(),
				mcp.Description("A list of file paths to process. The pipeline will be applied to each file."),
				mcp.Items(map[string]interface{}{"type": "string"}),
			),
			mcp.WithString("pipeline_definition",
				mcp.Required(),
				mcp.Description(`The JSON definition of the pipeline to run.
Example for converting to PDF:
{
  "processors": [
    {
      "actionconvert": true,
      "settingsconvert": { "saveformat": "pdf" }
    }
  ]
}
Example for adding a watermark:
{
  "processors": [
    {
      "actionwatermark": true,
      "settingswatermarks": [
        {
          "watermarktype": "Text",
          "text": "Confidential",
          "fontfamily": "HELVETICA",
          "fontsize": 72,
          "color": "red",
          "rotation": -45,
          "opacity": 0.5
        }
      ]
    }
  ]
}`),
			),
		),
		createPipelineHandler(cfg),
	)

	if err := server.ServeStdio(s); err != nil {
		log.Fatalf("Server error: %v\n", err)
	}
}

func createPipelineHandler(cfg *Config) server.ToolHandlerFunc {
	return func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		nc, err := nats.Connect(cfg.NatsURL)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to nats: %w", err)
		}
		defer nc.Close()

		js, err := nc.JetStream()
		if err != nil {
			return nil, fmt.Errorf("failed to get jetstream context: %w", err)
		}

		args, ok := req.Params.Arguments.(map[string]interface{})
		if !ok {
			return mcp.NewToolResultError("invalid arguments format"), nil
		}

		filesArg, _ := args["files"]
		files, _ := filesArg.([]interface{})
		pipelineDefStr, _ := args["pipeline_definition"].(string)

		pipeline, err := model.UnmarshalPipeline([]byte(pipelineDefStr))
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal pipeline definition: %w", err)
		}

		// Set license key in pipeline for server-side processing
		if cfg.IsLicensed {
			pipeline.Licencekey = &cfg.LicenseKey
		}

		pplID := uuid.New().String()
		pipeline.ID = &pplID

		kv, err := js.CreateKeyValue(&nats.KeyValueConfig{
			Bucket: GetPipelineStatusKVName(pplID),
			TTL:    1 * time.Hour,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create pipeline status kv store: %w", err)
		}

		inputArtifacts := make([]*model.Artifact, 0, len(files))
		for _, fileArg := range files {
			filePath, _ := fileArg.(string)
			artifact, err := uploadFileAndCreateArtifact(js, pplID, filePath)
			if err != nil {
				return nil, fmt.Errorf("failed to process file %s: %w", filePath, err)
			}
			inputArtifacts = append(inputArtifacts, artifact)
		}
		pipeline.Inputs = &model.Artifact{
			Children:   inputArtifacts,
			Clientinfo: &model.ClientInfo{Iscontainer: func(b bool) *bool { return &b }(true)},
		}

		streamName := GetPipelineRunName(natsServerName, "pipeline")
		if err := AddStream(js, streamName); err != nil {
			return nil, fmt.Errorf("failed to create nats stream: %w", err)
		}

		pipelineJSON, err := pipeline.Marshal()
		if err != nil {
			return nil, fmt.Errorf("failed to marshal pipeline: %w", err)
		}

		publishSubject := fmt.Sprintf("%s.%s", streamName, pplID)
		if _, err := js.Publish(publishSubject, pipelineJSON); err != nil {
			return nil, fmt.Errorf("failed to publish job: %w", err)
		}

		watcher, err := kv.Watch("status")
		if err != nil {
			return nil, fmt.Errorf("failed to watch pipeline status: %w", err)
		}
		defer watcher.Stop()

		timeout := time.After(2 * time.Minute)
		for {
			select {
			case entry := <-watcher.Updates():
				if entry != nil && string(entry.Value()) == "Completed" {
					goto COMPLETED
				}
				if entry != nil && string(entry.Value()) == "Errored" {
					return mcp.NewToolResultError("pipeline processing failed"), nil
				}
			case <-timeout:
				return nil, fmt.Errorf("timed out waiting for pipeline completion")
			}
		}

	COMPLETED:
		finalStateEntry, err := kv.Get("pipeline")
		if err != nil {
			return nil, fmt.Errorf("failed to get final pipeline state: %w", err)
		}

		var resultPipeline model.Pipeline
		if err := json.Unmarshal(finalStateEntry.Value(), &resultPipeline); err != nil {
			return nil, fmt.Errorf("failed to unmarshal final pipeline state: %w", err)
		}

		if resultPipeline.Output == nil || resultPipeline.Output.Asset == nil || resultPipeline.Output.Asset.Artifactnats == nil {
			return mcp.NewToolResultError("pipeline completed but returned no output artifact"), nil
		}

		// 6. Generate Secure Download Token and URL
		token := uuid.New().String()
		tokenMutex.Lock()
		tokenStore[token] = DownloadToken{
			Bucket:    *resultPipeline.Output.Asset.Artifactnats.Bucket,
			Key:       *resultPipeline.Output.Asset.Artifactnats.Natsid,
			ExpiresAt: time.Now().Add(5 * time.Minute),
		}
		tokenMutex.Unlock()

		downloadURL := fmt.Sprintf("http://localhost%s/download/%s", httpPort, token)
		return mcp.NewToolResultText(downloadURL), nil
	}
}

func uploadFileAndCreateArtifact(js nats.JetStreamContext, bucket, filePath string) (*model.Artifact, error) {
	store, err := js.ObjectStore(bucket)
	if err != nil {
		store, err = js.CreateObjectStore(&nats.ObjectStoreConfig{
			Bucket:  bucket,
			TTL:     24 * time.Hour,
			Storage: nats.FileStorage,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create object store bucket '%s': %w", bucket, err)
		}
	}

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	objectID := uuid.New().String()
	shortName := filepath.Base(filePath)
	mimetype := "application/octet-stream"

	_, err = store.PutBytes(objectID, fileContent)
	if err != nil {
		return nil, fmt.Errorf("failed to upload file to object store: %w", err)
	}

	storeTypeNats := model.Nats
	isContainerFalse := false

	artifact := &model.Artifact{
		ID: &objectID,
		Asset: &model.Asset{
			Artifactnats: &model.ArtifactNATS{
				Bucket: &bucket,
				Natsid: &objectID,
			},
			Storetype: &storeTypeNats,
			Mimetype:  &mimetype,
		},
		Clientinfo: &model.ClientInfo{
			Fullname:    &filePath,
			Shortname:   &shortName,
			Iscontainer: &isContainerFalse,
			Mimetype:    &mimetype,
		},
	}

	return artifact, nil
}

func startDownloadServer() {
	http.HandleFunc("/download/", downloadHandler)
	log.Printf("Starting download server on %s", httpPort)
	if err := http.ListenAndServe(httpPort, nil); err != nil {
		log.Fatalf("Failed to start download server: %v", err)
	}
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	token := filepath.Base(r.URL.Path)

	tokenMutex.Lock()
	defer tokenMutex.Unlock()

	download, ok := tokenStore[token]
	if !ok || time.Now().After(download.ExpiresAt) {
		http.Error(w, "Invalid or expired download token", http.StatusNotFound)
		return
	}

	// Invalidate token after first use
	delete(tokenStore, token)

	cfg := &Config{NatsURL: os.Getenv("NATS_URL")}
	if cfg.NatsURL == "" {
		cfg.NatsURL = nats.DefaultURL
	}

	nc, err := nats.Connect(cfg.NatsURL)
	if err != nil {
		http.Error(w, "Failed to connect to NATS", http.StatusInternalServerError)
		return
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		http.Error(w, "Failed to get JetStream context", http.StatusInternalServerError)
		return
	}

	store, err := js.ObjectStore(download.Bucket)
	if err != nil {
		http.Error(w, "Failed to access object store", http.StatusInternalServerError)
		return
	}

	obj, err := store.Get(download.Key)
	if err != nil {
		http.Error(w, "Failed to retrieve file", http.StatusInternalServerError)
		return
	}
	defer obj.Close()

	info, err := obj.Info()
	if err != nil {
		http.Error(w, "Failed to get file info", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+info.Name)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", info.Size))

	if _, err := io.Copy(w, obj); err != nil {
		log.Printf("Error streaming file to client: %v", err)
	}
}
