const os = require('os');

/**
 * Get platform-specific information for binary downloads
 */
function getPlatformInfo() {
  const platform = os.platform();
  const arch = os.arch();
  
  const platformMap = {
    'win32': 'Windows',
    'darwin': 'Darwin',
    'linux': 'Linux'
  };
  
  const archMap = {
    'x64': 'x86_64',
    'arm64': 'arm64',
    'arm': 'arm64'
  };
  
  return {
    platform: platformMap[platform],
    arch: archMap[arch],
    raw: { platform, arch },
    binaryName: platform === 'win32' ? 'df-mcp-server.exe' : 'df-mcp-server',
    archiveExtension: platform === 'win32' ? 'zip' : 'tar.gz'
  };
}

/**
 * Validate if the current platform is supported
 */
function isSupportedPlatform() {
  const { platform, arch } = getPlatformInfo();
  return platform && arch;
}

/**
 * Get the cache directory for storing binaries
 */
function getCacheDirectory() {
  const homeDir = os.homedir();
  const platform = os.platform();
  
  let cacheBase;
  if (process.env.XDG_CACHE_HOME) {
    cacheBase = process.env.XDG_CACHE_HOME;
  } else if (platform === 'win32') {
    cacheBase = process.env.LOCALAPPDATA || path.join(homeDir, 'AppData', 'Local');
  } else if (platform === 'darwin') {
    cacheBase = path.join(homeDir, 'Library', 'Caches');
  } else {
    cacheBase = path.join(homeDir, '.cache');
  }
  
  return path.join(cacheBase, 'df-mcp-server');
}

module.exports = {
  getPlatformInfo,
  isSupportedPlatform,
  getCacheDirectory
};
