# Makefile for ai-starter

.PHONY: build test lint format check clean install docs help

# Default target
.DEFAULT_GOAL := help

# Binary name
BINARY_NAME=ai

# Build the application
build: ## Build the application
	@echo "Building..."
	bun run build

# Build standalone binary
build-binary: ## Build standalone binary
	@echo "Building standalone binary..."
	bun run build:binary

# Run tests
test: ## Run tests
	@echo "Running tests..."
	bun test

# Run tests in watch mode
test-watch: ## Run tests in watch mode
	@echo "Running tests in watch mode..."
	bun test --watch

# Run linter
lint: ## Run linter
	@echo "Running linter..."
	bun run lint

# Format code
format: ## Format code
	@echo "Formatting code..."
	bun run format

# Check and fix code (lint + format)
check: ## Check and fix code (lint + format)
	@echo "Checking and fixing code..."
	bun run check

# Clean build artifacts
clean: ## Clean build artifacts
	@echo "Cleaning..."
	rm -f ${BINARY_NAME}
	rm -rf dist/
	rm -rf site/
	rm -rf node_modules/

# Install dependencies
install: ## Install dependencies
	@echo "Installing dependencies..."
	bun install

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

# Show help
help: ## Show this help message
	@echo "Available targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'
