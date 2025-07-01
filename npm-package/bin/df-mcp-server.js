#!/usr/bin/env node

const { spawn } = require('child_process');
const path = require('path');
const fs = require('fs').promises;
const os = require('os');
const https = require('https');
const { pipeline } = require('stream/promises');
const { createWriteStream, createReadStream } = require('fs');
const { promisify } = require('util');
const stream = require('stream');

// Package version (get from package.json)
const packageJson = require('../package.json');
const PACKAGE_VERSION = packageJson.version;

// GitHub repository details
const GITHUB_REPO = 'GlobalCents-DocumentFactory/mcp-documentfactory';

// Auto-detect platform and architecture
const platform = os.platform();
const arch = os.arch();

// Map to GoReleaser naming convention
const platformMap = {
  'win32': 'Windows',
  'darwin': 'Darwin',
  'linux': 'Linux'
};

const archMap = {
  'x64': 'x86_64',
  'arm64': 'arm64',
  'arm': 'arm64' // Fallback for some ARM systems
};

// Get the binary name for the current platform
function getBinaryName() {
  return platform === 'win32' ? 'df-mcp-server.exe' : 'df-mcp-server';
}

// Get the archive name for the current platform
function getArchiveName(version) {
  const platformName = platformMap[platform];
  const archName = archMap[arch];
  
  if (!platformName || !archName) {
    throw new Error(`Unsupported platform: ${platform}-${arch}`);
  }
  
  const extension = platform === 'win32' ? 'zip' : 'tar.gz';
  return `df-mcp-server_${platformName}_${archName}.${extension}`;
}

// Get the cache directory for storing binaries
function getCacheDir() {
  const homeDir = os.homedir();
  const cacheBase = process.env.XDG_CACHE_HOME || 
                   (platform === 'win32' ? path.join(homeDir, 'AppData', 'Local') : 
                    platform === 'darwin' ? path.join(homeDir, 'Library', 'Caches') : 
                    path.join(homeDir, '.cache'));
  
  return path.join(cacheBase, 'df-mcp-server');
}

// Download a file from URL
async function downloadFile(url, outputPath) {
  return new Promise((resolve, reject) => {
    const file = createWriteStream(outputPath);
    const request = https.get(url, (response) => {
      if (response.statusCode === 302 || response.statusCode === 301) {
        // Handle redirect
        return downloadFile(response.headers.location, outputPath).then(resolve).catch(reject);
      }
      
      if (response.statusCode !== 200) {
        reject(new Error(`Failed to download: ${response.statusCode} ${response.statusMessage}`));
        return;
      }
      
      response.pipe(file);
      
      file.on('finish', () => {
        file.close();
        resolve();
      });
      
      file.on('error', (err) => {
        fs.unlink(outputPath).catch(() => {}); // Clean up on error
        reject(err);
      });
    });
    
    request.on('error', (err) => {
      fs.unlink(outputPath).catch(() => {}); // Clean up on error
      reject(err);
    });
  });
}

// Extract archive (simple implementation for tar.gz and zip)
async function extractArchive(archivePath, extractDir, binaryName) {
  const { spawn } = require('child_process');
  
  return new Promise((resolve, reject) => {
    let extractCmd, extractArgs;
    
    if (archivePath.endsWith('.tar.gz')) {
      extractCmd = 'tar';
      extractArgs = ['-xzf', archivePath, '-C', extractDir];
    } else if (archivePath.endsWith('.zip')) {
      if (platform === 'win32') {
        // Use PowerShell on Windows
        extractCmd = 'powershell';
        extractArgs = ['-Command', `Expand-Archive -Path "${archivePath}" -DestinationPath "${extractDir}"`];
      } else {
        extractCmd = 'unzip';
        extractArgs = ['-q', archivePath, '-d', extractDir];
      }
    } else {
      reject(new Error(`Unsupported archive format: ${archivePath}`));
      return;
    }
    
    const process = spawn(extractCmd, extractArgs, { stdio: 'inherit' });
    
    process.on('close', (code) => {
      if (code === 0) {
        resolve();
      } else {
        reject(new Error(`Extraction failed with code ${code}`));
      }
    });
    
    process.on('error', reject);
  });
}

// Ensure binary exists and is up to date
async function ensureBinary() {
  const cacheDir = getCacheDir();
  const binaryName = getBinaryName();
  const binaryPath = path.join(cacheDir, binaryName);
  
  try {
    // Check if binary already exists
    await fs.access(binaryPath);
    
    // TODO: Add version check here if needed
    // For now, assume existing binary is correct version
    return binaryPath;
  } catch (error) {
    // Binary doesn't exist, need to download
  }
  
  // Create cache directory
  await fs.mkdir(cacheDir, { recursive: true });
  
  // Download and extract binary
  const archiveName = getArchiveName(PACKAGE_VERSION);
  const downloadUrl = `https://github.com/${GITHUB_REPO}/releases/download/v${PACKAGE_VERSION}/${archiveName}`;
  const archivePath = path.join(cacheDir, archiveName);
  
  console.error(`Downloading DocumentFactory MCP Server v${PACKAGE_VERSION}...`);
  console.error(`Platform: ${platform}-${arch}`);
  console.error(`URL: ${downloadUrl}`);
  
  try {
    await downloadFile(downloadUrl, archivePath);
    console.error('Download completed, extracting...');
    
    await extractArchive(archivePath, cacheDir, binaryName);
    console.error('Extraction completed');
    
    // Clean up archive
    await fs.unlink(archivePath);
    
    // Make binary executable on Unix systems
    if (platform !== 'win32') {
      const { chmod } = require('fs').promises;
      await chmod(binaryPath, 0o755);
    }
    
    console.error('DocumentFactory MCP Server ready!');
    return binaryPath;
    
  } catch (error) {
    console.error('Failed to download or extract binary:', error.message);
    console.error('');
    console.error('This might be due to:');
    console.error('1. Network connectivity issues');
    console.error('2. GitHub rate limiting');
    console.error('3. Missing release assets');
    console.error('4. Unsupported platform');
    console.error('');
    console.error('Please check:');
    console.error(`- Release exists: https://github.com/${GITHUB_REPO}/releases/tag/v${PACKAGE_VERSION}`);
    console.error(`- Platform support: ${platform}-${arch}`);
    process.exit(1);
  }
}

// Main function
async function main() {
  try {
    const binaryPath = await ensureBinary();
    
    // Set up environment variables
    const env = { ...process.env };
    
    // Use fixed NATS URL for DocumentFactory's public infrastructure
    env.NATS_URL = 'nats://nats.documentfactory.com:4222';
    
    // Pass through LICENCEKEY if provided
    if (process.env.LICENCEKEY) {
      env.LICENCEKEY = process.env.LICENCEKEY;
    }
    
    // Execute the binary with all arguments
    const child = spawn(binaryPath, process.argv.slice(2), {
      stdio: 'inherit',
      env: env
    });
    
    // Handle process exit
    child.on('exit', (code, signal) => {
      if (signal) {
        process.kill(process.pid, signal);
      } else {
        process.exit(code || 0);
      }
    });
    
    // Handle SIGINT (Ctrl+C) and SIGTERM
    process.on('SIGINT', () => {
      child.kill('SIGINT');
    });
    
    process.on('SIGTERM', () => {
      child.kill('SIGTERM');
    });
    
  } catch (error) {
    console.error('Error:', error.message);
    process.exit(1);
  }
}

// Handle unhandled rejections
process.on('unhandledRejection', (reason, promise) => {
  console.error('Unhandled Rejection at:', promise, 'reason:', reason);
  process.exit(1);
});

// Run the main function
main().catch((error) => {
  console.error('Fatal error:', error);
  process.exit(1);
});
