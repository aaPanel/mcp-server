# aaPanel MCP project Makefile

# Set variables
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod
BINARY_NAME=mcp-btpanel
BUILD_DIR=build

# System information
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

# Main file path
MAIN_PATH=main.go

# Default target
.PHONY: all
all: build

# Build application
.PHONY: build
build:
    @echo "Building project..."
    @mkdir -p $(BUILD_DIR)
    @$(GOBUILD) -trimpath -ldflags "-s -w" -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
    @echo "Build completed: $(BUILD_DIR)/$(BINARY_NAME)"

# Windows specific build
.PHONY: build-windows
build-windows:
    @echo "Building Windows version..."
    @mkdir -p $(BUILD_DIR)
    @GOOS=windows GOARCH=amd64 $(GOBUILD) -trimpath -ldflags "-s -w" -o $(BUILD_DIR)/$(BINARY_NAME).exe $(MAIN_PATH)
    @echo "Windows version built: $(BUILD_DIR)/$(BINARY_NAME).exe"

# Help information
.PHONY: help
help:
    @echo "Available commands:"
    @echo "  make build          - Build application"
    @echo "  make build-windows  - Build Windows version"
    @echo "  make clean          - Clean build artifacts"
    @echo "  make run            - Run application"
    @echo "  make test           - Run tests"
    @echo "  make tidy           - Tidy Go module dependencies"
    @echo "  make help           - Show help information"