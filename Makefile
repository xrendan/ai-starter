# Makefile for ai-starter

.PHONY: build test lint clean install docs help

# Default target
.DEFAULT_GOAL := help

# Binary name
BINARY_NAME=ai

# Build the application
build: ## Build the application
	@echo "Building..."
	go build -o ${BINARY_NAME} .

# Run tests
test: ## Run tests
	@echo "Running tests..."
	go test -v -race -coverprofile=coverage.txt ./...

# Run tests with coverage
coverage: test ## Run tests with coverage report
	@echo "Generating coverage report..."
	go tool cover -html=coverage.txt

# Run linter
lint: ## Run linter
	@echo "Running linter..."
	golangci-lint run

# Clean build artifacts
clean: ## Clean build artifacts
	@echo "Cleaning..."
	rm -f ${BINARY_NAME}
	rm -f coverage.txt
	rm -rf dist/
	rm -rf site/

# Install the binary
install: build ## Install the binary to $GOPATH/bin
	@echo "Installing..."
	go install .

# Serve documentation
docs: ## Serve documentation locally
	@echo "Starting documentation server..."
	mkdocs serve

# Build documentation
docs-build: ## Build documentation
	@echo "Building documentation..."
	mkdocs build

# Deploy documentation
docs-deploy: ## Deploy documentation to GitHub Pages
	@echo "Deploying documentation..."
	mkdocs gh-deploy

# Format code
fmt: ## Format code
	@echo "Formatting code..."
	go fmt ./...

# Run go mod tidy
tidy: ## Run go mod tidy
	@echo "Tidying dependencies..."
	go mod tidy

# Show help
help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'
