# DocumentFactory MCP Server

This is an MCP (Model Context Protocol) server for the DocumentFactory pipeline. It runs locally and connects to a NATS server to process document jobs, enabling seamless integration with AI tools and workflows.

## What is this?

The DocumentFactory MCP Server allows AI assistants and other MCP-compatible clients to interact with your document processing pipeline. It provides a bridge between AI tools (like Claude, ChatGPT, or VS Code extensions) and your DocumentFactory infrastructure via NATS messaging.

## Installation

The easiest way to use the DocumentFactory MCP Server is with NPX:

```bash
npx @documentfactory/mcp-server
```

This automatically downloads and runs the correct binary for your platform - no manual installation required!

## Configuration

### License Key Setup

The server requires a license key to differentiate between demo and licensed usage:

- **Demo Mode**: No license key required. Converted PDFs will include a watermark.
- **Licensed Mode**: Provide a valid license key to remove watermarks from converted PDFs.

**Setting up the License Key:**

**On Linux/macOS:**
```bash
LICENCEKEY="your-license-key" npx @documentfactory/mcp-server
```

**On Windows (Command Prompt):**
```cmd
set LICENCEKEY=your-license-key && npx @documentfactory/mcp-server
```

**On Windows (PowerShell):**
```powershell
$env:LICENCEKEY="your-license-key"; npx @documentfactory/mcp-server
```

If no license key is provided, the server will run in demo mode with watermarked output.

## Usage

This server is designed to be run in `stdio` mode by MCP clients. It's not meant to be run directly by end users, but rather integrated into AI tools and applications.

### Running the Server

**Demo Mode (with watermarks):**
```bash
npx @documentfactory/mcp-server
```

**Licensed Mode (no watermarks):**
```bash
LICENCEKEY="your-license-key" npx @documentfactory/mcp-server
```

The server will start and wait for MCP client connections via standard input/output.

### Integration with AI Tools

To use this server with AI tools:

1. **VS Code with MCP Extension**: Add the server configuration to your MCP settings
2. **Claude Desktop**: Add the server to your Claude configuration
3. **Other MCP Clients**: Follow your client's documentation for adding MCP servers

**Example configuration for demo mode (with watermarks):**
```json
{
  "mcpServers": {
    "documentfactory": {
      "command": "npx",
      "args": ["@documentfactory/mcp-server"]
    }
  }
}
```

**Example configuration for licensed mode (no watermarks):**
```json
{
  "mcpServers": {
    "documentfactory": {
      "command": "npx",
      "args": ["@documentfactory/mcp-server"],
      "env": {
        "LICENCEKEY": "your-license-key"
      }
    }
  }
}
```

## Features

- **Document Processing**: Send documents through your DocumentFactory pipeline
- **Demo & Licensed Modes**: Try the service with watermarked PDFs or use with a license for clean output
- **NATS Integration**: Built-in connection to DocumentFactory's messaging infrastructure
- **Cross-Platform**: Works on Windows, macOS, and Linux
- **MCP Compatible**: Works with any MCP-compatible AI tool or client
- **Zero Installation**: Run instantly with NPX - no downloads or setup required

## Troubleshooting

### Common Issues

**Server won't start:**
- Ensure you have Node.js installed for NPX to work
- Check your internet connection for downloading the binary
- Verify the license key format if using licensed mode

**Watermarks appearing on PDFs:**
- This is normal behavior in demo mode
- Provide a valid `LICENCEKEY` environment variable to remove watermarks

**Connection issues:**
- The server connects to DocumentFactory's public infrastructure automatically
- Ensure your firewall allows outbound connections
- Check that NPX can download packages from the internet

**NPX issues:**
```bash
# Clear NPX cache if having download issues
npx clear-npx-cache
```

### Getting Help

If you encounter issues:

1. Try running in demo mode first (without license key) to verify basic functionality
2. Check that NPX and Node.js are properly installed
3. Verify your internet connection for package downloads
4. Ensure your license key is valid if using licensed mode
5. Check the server logs for error messages

For license key issues or technical support, please contact DocumentFactory support.

## Development

### Building from Source

If you need to build from source:

```bash
git clone https://github.com/GlobalCents-DocumentFactory/mcp-documentfactory.git
cd mcp-documentfactory
make build
```

### Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests: `make test-all`
5. Submit a pull request

## System Requirements

- **Operating System**: Windows 10+, macOS 10.15+, or Linux (any modern distribution)
- **Node.js**: Version 16+ (required for NPX)
- **Network**: Internet access for package downloads and DocumentFactory service connection
- **Memory**: Minimal (typically < 50MB RAM)
- **Disk Space**: Minimal (NPX handles caching automatically)

## License

This project is licensed under the terms specified in the LICENSE file.

## Support

For technical support and documentation, please refer to the project's GitHub repository or contact your system administrator.
