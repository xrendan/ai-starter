package setup

import (
	"fmt"
	"os"
	"path/filepath"
)

// SetupProjectStructure creates standard project directories and files
func SetupProjectStructure(projectName string) error {
	// Create standard directories
	dirs := []string{
		"cmd",
		"internal",
		"pkg",
		"test",
		"scripts",
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create %s: %w", dir, err)
		}
	}

	// Create .gitignore
	gitignore := `# Binaries
*.exe
*.exe~
*.dll
*.so
*.dylib
ai
dist/
build/

# Test binary, built with 'go test -c'
*.test

# Output of the go coverage tool
*.out
coverage.txt

# Go workspace file
go.work

# Dependencies
vendor/

# IDE
.idea/
.vscode/
*.swp
*.swo
*~

# OS
.DS_Store
Thumbs.db

# Documentation build
site/

# Temporary files
tmp/
temp/
*.tmp

# Environment
.env
.env.local

# AI tools
.cursor/
`

	if err := os.WriteFile(".gitignore", []byte(gitignore), 0644); err != nil {
		return fmt.Errorf("failed to create .gitignore: %w", err)
	}

	// Create Makefile
	makefile := `# Makefile for ai-starter

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
`

	if err := os.WriteFile("Makefile", []byte(makefile), 0644); err != nil {
		return fmt.Errorf("failed to create Makefile: %w", err)
	}

	// Create .golangci.yml for linting
	golangciConfig := `run:
  timeout: 5m
  tests: true

linters:
  enable:
    - errcheck
    - gosimple
    - govet
    - ineffassign
    - staticcheck
    - unused
    - gofmt
    - goimports
    - misspell
    - unconvert
    - unparam
    - gosec
    - gocritic

linters-settings:
  errcheck:
    check-type-assertions: true
    check-blank: true
  govet:
    check-shadowing: true
  goimports:
    local-prefixes: github.com/xrendan/ai-starter

issues:
  exclude-use-default: false
  max-issues-per-linter: 0
  max-same-issues: 0
`

	if err := os.WriteFile(".golangci.yml", []byte(golangciConfig), 0644); err != nil {
		return fmt.Errorf("failed to create .golangci.yml: %w", err)
	}

	// Create CONTRIBUTING.md
	contributing := `# Contributing to AI-Starter

Thank you for your interest in contributing! This document provides guidelines for contributing to this project.

## Perfect Commit Structure

This project follows Simon Willison's "perfect commit" philosophy. Each commit should include:

1. **Implementation** - A single, focused change
2. **Tests** - Demonstrate the change works
3. **Documentation** - Keep docs in sync
4. **Issue Link** - Provide context

See [docs/development/contributing.md](docs/development/contributing.md) for detailed guidelines.

## Getting Started

1. Fork the repository
2. Clone your fork
3. Create a feature branch
4. Make your changes following the perfect commit structure
5. Submit a pull request

## Development Workflow

1. Create an issue describing the change
2. Reference the issue in your commits
3. Include tests for new functionality
4. Update documentation in the same commit
5. Ensure CI passes

## Code of Conduct

Be respectful and constructive. We're all here to build great software together.

## Questions?

Open an issue or start a discussion!
`

	if err := os.WriteFile("CONTRIBUTING.md", []byte(contributing), 0644); err != nil {
		return fmt.Errorf("failed to create CONTRIBUTING.md: %w", err)
	}

	// Create sample test file
	testExample := `package main

import (
	"testing"
)

func TestPlaceholder(t *testing.T) {
	// This is a placeholder test to establish testing infrastructure
	// Remove this when you add real tests

	t.Run("example test", func(t *testing.T) {
		if 1+1 != 2 {
			t.Error("math is broken")
		}
	})
}
`

	testDir := "test"
	if err := os.MkdirAll(testDir, 0755); err != nil {
		return fmt.Errorf("failed to create test directory: %w", err)
	}

	testPath := filepath.Join(testDir, "example_test.go")
	if err := os.WriteFile(testPath, []byte(testExample), 0644); err != nil {
		return fmt.Errorf("failed to create example test: %w", err)
	}

	fmt.Println("✓ Project structure created")
	return nil
}

// CreateReadme creates a comprehensive README
func CreateReadme(projectName string) error {
	readme := fmt.Sprintf(`# %s

AI-driven repository setup and management tool.

## Overview

This tool helps set up repositories optimized for AI-driven development workflows. It enforces best practices including:

- 📚 **Documentation-First Development** - Using MkDocs with Material theme
- 🎯 **Perfect Commits** - Following Simon Willison's commit structure
- 🤖 **AI Prompt Tracking** - Using git-ai to track prompts as source code
- ✅ **Automated Testing** - Comprehensive test coverage required
- 🔄 **CI/CD** - GitHub Actions for validation and deployment

## Installation

### From Source

\'\'\'bash
git clone https://github.com/xrendan/%s.git
cd %s
go build -o ai .
sudo mv ai /usr/local/bin/
\'\'\'

### From Release

Download the latest release for your platform from the [releases page](https://github.com/xrendan/%s/releases).

## Quick Start

Initialize a new AI-driven repository:

\'\'\'bash
ai init my-project
cd my-project
\'\'\'

This will set up:
- Documentation structure with MkDocs
- Git commit templates and hooks
- git-ai for prompt tracking
- GitHub Actions workflows
- Testing infrastructure

## Perfect Commit Structure

Each commit should include:

1. **Implementation** - A single, focused change
2. **Tests** - Demonstrate the change works
3. **Documentation** - Keep docs in sync with code
4. **Issue Link** - Provide context and rationale

### Example Commit

\'\'\'
feat: Add user authentication (#42)

Implementation:
- Added JWT-based authentication middleware
- Implemented login and logout endpoints

Tests:
- Added auth middleware tests
- Added integration tests for login flow

Documentation:
- Updated API documentation with auth endpoints
- Added authentication guide to docs/

Closes #42
\'\'\'

## AI Prompt Tracking

This project uses [git-ai](https://github.com/acunniffe/git-ai) to track AI-generated code and prompts.

Prompts are stored in git notes and tracked alongside code changes. This ensures:
- Full visibility into AI contributions
- Prompts are preserved for future reference
- Understanding of code evolution and decisions

## Documentation

Documentation is built with MkDocs and deployed to GitHub Pages.

### Local Development

\'\'\'bash
# Install dependencies
pip install -r docs/requirements.txt

# Serve locally
mkdocs serve

# Visit http://localhost:8000
\'\'\'

### Structure

- \'docs/\' - Main documentation
- \'docs/adr/\' - Architectural Decision Records
- \'docs/guides/\' - User guides and tutorials
- \'docs/prompts/\' - AI prompts and context
- \'docs/development/\' - Development documentation

## Development

### Prerequisites

- Go 1.21+
- Git
- Python 3.8+ (for docs)
- Make (optional)

### Building

\'\'\'bash
make build
\'\'\'

### Testing

\'\'\'bash
make test
\'\'\'

### Linting

\'\'\'bash
make lint
\'\'\'

## CI/CD

GitHub Actions workflows handle:

- **CI** - Run tests, linting, and builds on every push
- **Documentation** - Validate and deploy docs on changes
- **Release** - Build multi-platform binaries on tags

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for detailed contribution guidelines.

All contributions must follow the perfect commit structure.

## License

See [LICENSE](LICENSE) for details.

## Acknowledgments

- [Simon Willison](https://simonwillison.net/) - Perfect commit philosophy
- [git-ai](https://github.com/acunniffe/git-ai) - AI code tracking
- [MkDocs](https://www.mkdocs.org/) - Documentation framework

## Support

- 📖 [Documentation](https://xrendan.github.io/%s/)
- 🐛 [Issue Tracker](https://github.com/xrendan/%s/issues)
- 💬 [Discussions](https://github.com/xrendan/%s/discussions)
`, projectName, projectName, projectName, projectName, projectName, projectName, projectName)

	if err := os.WriteFile("README.md", []byte(readme), 0644); err != nil {
		return fmt.Errorf("failed to create README.md: %w", err)
	}

	fmt.Println("✓ README created")
	return nil
}
