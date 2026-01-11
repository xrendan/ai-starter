package setup

import (
	"fmt"
	"os"
	"path/filepath"
)

// SetupDocs creates the documentation structure with MkDocs
func SetupDocs(projectName string) error {
	// Create docs directory structure
	dirs := []string{
		"docs",
		"docs/adr",           // Architectural Decision Records
		"docs/api",           // API documentation
		"docs/guides",        // User guides
		"docs/prompts",       // AI prompts and context
		"docs/development",   // Development documentation
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create %s: %w", dir, err)
		}
	}

	// Create mkdocs.yml configuration
	mkdocsConfig := fmt.Sprintf(`site_name: %s
site_description: AI-driven repository for %s
repo_url: https://github.com/xrendan/%s

theme:
  name: material
  palette:
    - scheme: default
      primary: indigo
      accent: indigo
      toggle:
        icon: material/brightness-7
        name: Switch to dark mode
    - scheme: slate
      primary: indigo
      accent: indigo
      toggle:
        icon: material/brightness-4
        name: Switch to light mode
  features:
    - navigation.tabs
    - navigation.sections
    - navigation.expand
    - search.suggest
    - search.highlight
    - content.code.copy

nav:
  - Home: index.md
  - Architecture:
      - Overview: adr/index.md
  - Development:
      - Setup: development/setup.md
      - Contributing: development/contributing.md
      - Testing: development/testing.md
  - Guides: guides/index.md
  - AI Prompts: prompts/index.md

markdown_extensions:
  - admonition
  - codehilite
  - toc:
      permalink: true
  - pymdownx.highlight:
      anchor_linenums: true
  - pymdownx.inlinehilite
  - pymdownx.snippets
  - pymdownx.superfences
  - pymdownx.details
  - pymdownx.tabbed:
      alternate_style: true

plugins:
  - search
  - git-revision-date-localized:
      enable_creation_date: true
`, projectName, projectName, projectName)

	if err := os.WriteFile("mkdocs.yml", []byte(mkdocsConfig), 0644); err != nil {
		return fmt.Errorf("failed to create mkdocs.yml: %w", err)
	}

	// Create index page
	indexContent := fmt.Sprintf(`# %s

Welcome to the %s documentation.

## Overview

This repository follows AI-driven development best practices, including:

- **Documentation-First**: All features start with documentation
- **Perfect Commits**: Following Simon Willison's commit structure
- **Prompt Tracking**: AI prompts tracked with git-ai
- **Automated Testing**: Comprehensive test coverage
- **CI/CD**: Automated validation and deployment

## Quick Start

See the [Setup Guide](development/setup.md) to get started.

## Architecture

Review our [Architectural Decision Records](adr/index.md) to understand key design decisions.

## Contributing

Please read our [Contributing Guide](development/contributing.md) before submitting changes.
`, projectName, projectName)

	if err := os.WriteFile("docs/index.md", []byte(indexContent), 0644); err != nil {
		return fmt.Errorf("failed to create docs/index.md: %w", err)
	}

	// Create ADR index
	adrIndex := `# Architectural Decision Records

This directory contains records of architectural decisions made in this project.

## What is an ADR?

An Architectural Decision Record (ADR) captures an important architectural decision along with its context and consequences.

## Format

Each ADR should include:

1. **Title**: A short, descriptive title
2. **Status**: Proposed, Accepted, Deprecated, or Superseded
3. **Context**: The issue motivating this decision
4. **Decision**: The change being proposed or decided
5. **Consequences**: The resulting context after applying the decision

## ADRs

- [ADR-001: Use MkDocs for Documentation](001-use-mkdocs.md)
- [ADR-002: Enforce Perfect Commit Structure](002-perfect-commits.md)
- [ADR-003: Track AI Prompts with git-ai](003-track-ai-prompts.md)
`

	if err := os.WriteFile("docs/adr/index.md", []byte(adrIndex), 0644); err != nil {
		return fmt.Errorf("failed to create ADR index: %w", err)
	}

	// Create sample ADRs
	if err := createSampleADRs(); err != nil {
		return err
	}

	// Create other documentation files
	if err := createDevelopmentDocs(projectName); err != nil {
		return err
	}

	// Create prompts documentation
	promptsIndex := `# AI Prompts

This directory contains important AI prompts and context used in developing this project.

## Why Track Prompts?

AI prompts are the new source code. Just as we track code changes, we should track the prompts and context that generate code.

## Organization

- **Context**: Important context and background information
- **Templates**: Reusable prompt templates
- **Examples**: Example prompts and their results

## Best Practices

1. Document prompts that generate significant code
2. Include context about why specific approaches were chosen
3. Track iterations and refinements
4. Link prompts to their corresponding commits
`

	if err := os.WriteFile("docs/prompts/index.md", []byte(promptsIndex), 0644); err != nil {
		return fmt.Errorf("failed to create prompts index: %w", err)
	}

	// Create guides index
	guidesIndex := `# Guides

User guides and tutorials for this project.

## Available Guides

Add your guides here as you create them.
`

	if err := os.WriteFile("docs/guides/index.md", []byte(guidesIndex), 0644); err != nil {
		return fmt.Errorf("failed to create guides index: %w", err)
	}

	// Create requirements.txt for MkDocs
	requirementsContent := `mkdocs>=1.5.0
mkdocs-material>=9.0.0
mkdocs-git-revision-date-localized-plugin>=1.2.0
`

	if err := os.WriteFile("docs/requirements.txt", []byte(requirementsContent), 0644); err != nil {
		return fmt.Errorf("failed to create requirements.txt: %w", err)
	}

	fmt.Println("✓ Documentation structure created")
	return nil
}

func createSampleADRs() error {
	adr001 := `# ADR-001: Use MkDocs for Documentation

**Status**: Accepted

**Date**: 2026-01-11

## Context

We need a documentation system that:
- Lives in the repository with the code
- Supports versioning alongside code changes
- Is easy to write and maintain
- Generates professional-looking documentation
- Integrates well with CI/CD

## Decision

We will use MkDocs with the Material theme for all project documentation.

## Consequences

**Positive**:
- Documentation lives in the repository
- Markdown is easy to write and review
- Material theme provides excellent UX
- Easy to deploy to GitHub Pages
- Can validate documentation in CI/CD

**Negative**:
- Requires Python for local preview
- Team needs to learn MkDocs configuration

## Alternatives Considered

- GitBook: More features but external hosting
- Jekyll: More complex configuration
- Docusaurus: React-based, heavier weight
`

	adr002 := `# ADR-002: Enforce Perfect Commit Structure

**Status**: Accepted

**Date**: 2026-01-11

## Context

Following Simon Willison's "perfect commit" philosophy, we want each commit to be:
1. A single, focused change (implementation)
2. Accompanied by tests
3. Accompanied by documentation
4. Linked to an issue for context

This makes commits more reviewable, understandable, and maintainable.

## Decision

We will enforce commit structure through:
- Commit message templates
- Git hooks that validate commit messages
- GitHub Actions that verify commits include tests and docs
- Required issue references

## Consequences

**Positive**:
- More maintainable git history
- Easier code review
- Better long-term understanding of changes
- Self-documenting codebase

**Negative**:
- Requires discipline and training
- May slow down initial development
- Not every commit needs to be perfect (typos, etc.)

## Implementation

- Commit template in .gitmessage
- commit-msg hook for validation
- CI checks for test and doc changes
`

	adr003 := `# ADR-003: Track AI Prompts with git-ai

**Status**: Accepted

**Date**: 2026-01-11

## Context

AI-generated code is becoming a significant portion of modern codebases. We need to:
- Track which code was AI-generated
- Preserve the prompts that generated code
- Understand the evolution of AI interactions
- Maintain accountability and reviewability

Prompts are the new source code - they should be versioned and tracked.

## Decision

We will use git-ai to:
- Automatically track AI-generated code
- Store prompts in git notes
- Score commits by AI contribution
- Preserve AI authorship through git operations

## Consequences

**Positive**:
- Full visibility into AI contributions
- Prompts are preserved for future reference
- Can analyze AI impact on codebase
- Supports multiple AI coding agents

**Negative**:
- Additional tool to install and maintain
- Git notes add some complexity
- Team needs to understand git-ai

## Alternatives Considered

- Manual prompt documentation: Too error-prone
- Comments in code: Clutters codebase
- Separate prompt repository: Loses connection to code
`

	files := map[string]string{
		"docs/adr/001-use-mkdocs.md":       adr001,
		"docs/adr/002-perfect-commits.md":  adr002,
		"docs/adr/003-track-ai-prompts.md": adr003,
	}

	for path, content := range files {
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to create %s: %w", path, err)
		}
	}

	return nil
}

func createDevelopmentDocs(projectName string) error {
	setupGuide := fmt.Sprintf(`# Setup Guide

## Prerequisites

- Go 1.21 or higher
- Git
- Python 3.8+ (for documentation)
- Make (optional)

## Installation

1. Clone the repository:
   \'\'\'bash
   git clone https://github.com/xrendan/%s.git
   cd %s
   \'\'\'

2. Install dependencies:
   \'\'\'bash
   go mod download
   \'\'\'

3. Install git-ai for prompt tracking:
   \'\'\'bash
   curl -sSL https://usegitai.com/install.sh | bash
   \'\'\'

4. Install documentation tools:
   \'\'\'bash
   pip install -r docs/requirements.txt
   \'\'\'

## Development

### Building

\'\'\'bash
go build -o ai .
\'\'\'

### Testing

\'\'\'bash
go test ./...
\'\'\'

### Documentation

Preview documentation locally:

\'\'\'bash
mkdocs serve
\'\'\'

Then visit http://localhost:8000

## Commit Guidelines

This project follows the "perfect commit" structure:

1. Each commit should be a focused, atomic change
2. Include tests for new functionality
3. Update documentation in the same commit
4. Reference the issue number in the commit message

Use the commit template:

\'\'\'bash
git config commit.template .gitmessage
\'\'\'
`, projectName, projectName)

	contributingGuide := `# Contributing Guide

Thank you for contributing! This guide will help you make effective contributions.

## Perfect Commit Structure

Following Simon Willison's philosophy, each commit should include:

### 1. Implementation
- A single, focused change
- Should be atomic and deployable
- Easy to review and understand

### 2. Tests
- Demonstrate the change works
- Prevent regressions
- Document expected behavior

### 3. Documentation
- Update relevant docs in the same commit
- Include API docs, user guides, or ADRs as needed
- Keep docs in sync with code

### 4. Issue Link
- Reference the issue number
- Provides context and discussion history
- Explains the "why" behind the change

## Workflow

1. **Create an issue** describing the change
2. **Create a branch** from main
3. **Make your changes** following the commit structure
4. **Run tests** to ensure everything works
5. **Update documentation** in the same commit
6. **Submit a PR** referencing the issue
7. **Address review feedback**

## Commit Message Format

\'\'\'
[Type] Brief description (#issue)

Implementation:
- What changed

Tests:
- What tests were added/modified

Documentation:
- What docs were updated

Closes #issue
\'\'\'

## Types

- **feat**: New feature
- **fix**: Bug fix
- **docs**: Documentation only
- **refactor**: Code refactoring
- **test**: Adding tests
- **chore**: Maintenance tasks

## Code Review

All changes require:
- Passing tests
- Updated documentation
- Code review approval
- CI/CD checks passing
`

	testingGuide := `# Testing Guide

## Test Philosophy

- Tests are part of every commit
- Tests document expected behavior
- Tests prevent regressions
- Tests enable confident refactoring

## Running Tests

Run all tests:
\'\'\'bash
go test ./...
\'\'\'

Run with coverage:
\'\'\'bash
go test -cover ./...
\'\'\'

Run verbose:
\'\'\'bash
go test -v ./...
\'\'\'

## Writing Tests

### Unit Tests

- Test individual functions and methods
- Use table-driven tests for multiple cases
- Mock external dependencies

### Integration Tests

- Test component interactions
- Use real dependencies when feasible
- Test error handling

### Example

\'\'\'go
func TestExample(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        want    string
        wantErr bool
    }{
        {
            name:  "valid input",
            input: "test",
            want:  "result",
        },
        // Add more cases
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Example(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("unexpected error: %v", err)
            }
            if got != tt.want {
                t.Errorf("got %v, want %v", got, tt.want)
            }
        })
    }
}
\'\'\'

## CI/CD

Tests run automatically on:
- Every push
- Every pull request
- Before deployment

All tests must pass before merging.
`

	files := map[string]string{
		"docs/development/setup.md":        setupGuide,
		"docs/development/contributing.md": contributingGuide,
		"docs/development/testing.md":      testingGuide,
	}

	for path, content := range files {
		// Create directory if it doesn't exist
		dir := filepath.Dir(path)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}

		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return fmt.Errorf("failed to create %s: %w", path, err)
		}
	}

	return nil
}
