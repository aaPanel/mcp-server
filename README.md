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
git clone https://github.com/yourusername/mcp_btpanel.git
cd mcp_btpanel

# Install dependencies
go mod tidy

# Build project
make build

# Build for Windows
.\build.bat build
```