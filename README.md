# aaPanel MCP Interface
## Features

- Get panel system information and network status
- Query PHP website list
- Create new PHP website
- Query MySQL database list
- Docker container management
  - View container list
  - View container details
- Docker image management
  - View local image list
- Email management
  - Add email account
  - View email list
- Get panel public configuration information
- More features under development...

## Requirements
- Go 1.18+
- aaPanel API access
- aaPanel API token

## Installation
### From source code

```bash
# Clone repository
git clone https://github.com/aaPanel/mcp-server.git
cd mcp_btpanel

# Install dependencies
go mod tidy

# Build project
make build

# Build for Windows
.\build.bat build
```

### Direct download
You can also download pre-compiled binaries from the Releases page .

## Configuration
### Environment variables
Configure the program through environment variables:
```bash
# Set aaPanel address
export BT_BASE_URL="http://your-panel-address:8888"

# Set aaPanel API token
export BT_API_TOKEN="your-api-token"
```

### Cursor configuration
When using in Cursor, configure through the following steps:

1. Open Cursor Settings > Extensions > MCP Tools
2. Add new MCP tool
3. Fill in the configuration in the following format:
```json
{
    "mcpServers": {
        "mcp-aapanel": {
            "command": "C:\\path\\to\\mcp-server.exe",
            "env": {
                "BT_BASE_URL": "http://192.168.xx.xx:8888/",
                "BT_API_TOKEN": "xxxxxxxxxxxxxxxxxxxxxxxx"
            }
        }
    }
}
```

### Adding new features
1. Create or modify files in the corresponding module directory
2. Define new tool constants and tool objects
3. Implement handler functions
4. Register tools in the registerTools function in main.go
## Build & Deployment
Build using Makefile:
```bash
# Build project
make build

# Build Windows version
make build-windows

# Clean build artifacts
make clean

# View more commands
make help
```

## License
This project is licensed under the MIT License. See the LICENSE file for details.

## Contributing
Welcome to submit issues and feature requests! If you want to contribute, please:

1. Fork this repository
2. Create your feature branch ( git checkout -b feature/amazing-feature )
3. Commit your changes ( git commit -m 'Add some amazing feature' )
4. Push to the branch ( git push origin feature/amazing-feature )
5. Open a Pull Request