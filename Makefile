.PHONY: build test clean install-npm test-npm publish-npm help

# Go binary name
BINARY_NAME=df-mcp-server
BUILD_DIR=./cmd/df-mcp-server

# Default target
help:
	@echo "DocumentFactory MCP Server - Development Commands"
	@echo "================================================"
	@echo ""
	@echo "Go Commands:"
	@echo "  build      - Build the Go binary"
	@echo "  test       - Run Go tests"
	@echo "  clean      - Clean build artifacts"
	@echo ""
	@echo "NPM Commands:"
	@echo "  install-npm - Install NPM package dependencies"
	@echo "  test-npm    - Test NPM package"
	@echo "  publish-npm - Publish NPM package (requires npm login)"
	@echo ""
	@echo "Release Commands:"
	@echo "  release     - Create a new release (requires tag)"
	@echo ""
	@echo "Environment Variables:"
	@echo "  LICENCEKEY  - License key for testing licensed mode"
	@echo "  NATS_URL    - NATS server URL (defaults to public server)"

# Build the Go binary
build:
	@echo "Building $(BINARY_NAME)..."
	go build -o $(BINARY_NAME) $(BUILD_DIR)
	@echo "✓ Build complete: ./$(BINARY_NAME)"

# Run Go tests
test:
	@echo "Running Go tests..."
	go test -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -f $(BINARY_NAME)
	rm -rf dist/
	@echo "✓ Clean complete"

# Install NPM package dependencies
install-npm:
	@echo "Installing NPM package dependencies..."
	cd npm-package && npm install
	@echo "✓ NPM dependencies installed"

# Test NPM package
test-npm: install-npm
	@echo "Testing NPM package..."
	./test-npm-package.sh
	@echo "✓ NPM package tests complete"

# Publish NPM package (must be logged in to npm)
publish-npm: install-npm
	@echo "Publishing NPM package..."
	@echo "Make sure you're logged in: npm login"
	@read -p "Press Enter to continue or Ctrl+C to cancel..."
	cd npm-package && npm publish --access public
	@echo "✓ NPM package published"

# Create a release using GoReleaser (requires tag)
release:
	@echo "Creating release with GoReleaser..."
	@if [ -z "$(shell git tag --points-at HEAD)" ]; then \
		echo "Error: No tag found at HEAD. Create a tag first:"; \
		echo "  git tag v1.0.0"; \
		echo "  git push origin v1.0.0"; \
		exit 1; \
	fi
	goreleaser release --clean
	@echo "✓ Release complete"

# Development setup
dev-setup:
	@echo "Setting up development environment..."
	go mod tidy
	go mod download
	cd npm-package && npm install
	@echo "✓ Development environment ready"

# Test both Go and NPM in sequence
test-all: test test-npm
	@echo "✓ All tests complete"

# Build and test everything
ci: build test test-npm
	@echo "✓ CI pipeline complete"
