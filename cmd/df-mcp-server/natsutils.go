package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func GetPipelineRunName(serverid, pipelineType string) string {
	return fmt.Sprintf("pplrun_%s_%s", serverid, pipelineType)
}

func GetPipelineStatusKVName(pipelineid string) string {
	return fmt.Sprintf("pplstatus_%s", pipelineid)
}

func AddStream(js nats.JetStreamContext, name string) error {
	conf := &nats.StreamConfig{
		Name:       name,
		Subjects:   []string{name + ".>"},
		Storage:    nats.FileStorage,
		Retention:  nats.WorkQueuePolicy,
		MaxBytes:   2 * 1024 * 1024 * 1024, // 2GB
		MaxMsgSize: 200 * 1024 * 1024,      // 200MB
	}
	_, err := js.AddStream(conf)

	if err != nil {
		if _, updateErr := js.UpdateStream(conf); updateErr != nil {
			return updateErr
		}
	}

	log.Printf("Initialized NATS Stream: %s", name)
	return nil
}
