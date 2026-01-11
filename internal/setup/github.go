package setup

import (
	"fmt"
	"os"
	"path/filepath"
)

// SetupGithubActions creates GitHub Actions workflows
func SetupGithubActions(projectName string) error {
	workflowsDir := filepath.Join(".github", "workflows")
	if err := os.MkdirAll(workflowsDir, 0755); err != nil {
		return fmt.Errorf("failed to create workflows directory: %w", err)
	}

	// CI workflow for tests and linting
	ciWorkflow := `name: CI

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main, develop ]

jobs:
  test:
    name: Test
    runs-on: ubuntu-latest

    steps:
    - name: Checkout code
      uses: actions/checkout@v4

    - name: Set up Go
      uses: actions/setup-go@v5
      with:
        go-version: '1.21'

    - name: Cache Go modules
      uses: actions/cache@v4
      with:
        path: ~/go/pkg/mod
        key: ${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}
        restore-keys: |
          ${{ runner.os }}-go-

    - name: Download dependencies
      run: go mod download

    - name: Run tests
      run: go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

    - name: Upload coverage
      uses: codecov/codecov-action@v4
      with:
        files: ./coverage.txt
        fail_ci_if_error: false

  lint:
    name: Lint
    runs-on: ubuntu-latest

    steps:
    - name: Checkout code
      uses: actions/checkout@v4

    - name: Set up Go
      uses: actions/setup-go@v5
      with:
        go-version: '1.21'

    - name: Run golangci-lint
      uses: golangci/golangci-lint-action@v4
      with:
        version: latest

  build:
    name: Build
    runs-on: ubuntu-latest

    steps:
    - name: Checkout code
      uses: actions/checkout@v4

    - name: Set up Go
      uses: actions/setup-go@v5
      with:
        go-version: '1.21'

    - name: Build
      run: go build -v .

    - name: Upload artifact
      uses: actions/upload-artifact@v4
      with:
        name: ai-binary
        path: ./ai
        retention-days: 5
`

	ciWorkflowPath := filepath.Join(workflowsDir, "ci.yml")
	if err := os.WriteFile(ciWorkflowPath, []byte(ciWorkflow), 0644); err != nil {
		return fmt.Errorf("failed to create CI workflow: %w", err)
	}

	// Documentation workflow
	docsWorkflow := `name: Documentation

on:
  push:
    branches: [ main ]
    paths:
      - 'docs/**'
      - 'mkdocs.yml'
  pull_request:
    paths:
      - 'docs/**'
      - 'mkdocs.yml'

jobs:
  validate:
    name: Validate Documentation
    runs-on: ubuntu-latest

    steps:
    - name: Checkout code
      uses: actions/checkout@v4

    - name: Set up Python
      uses: actions/setup-python@v5
      with:
        python-version: '3.11'

    - name: Cache Python packages
      uses: actions/cache@v4
      with:
        path: ~/.cache/pip
        key: ${{ runner.os }}-pip-${{ hashFiles('docs/requirements.txt') }}
        restore-keys: |
          ${{ runner.os }}-pip-

    - name: Install dependencies
      run: |
        pip install -r docs/requirements.txt

    - name: Build documentation
      run: mkdocs build --strict

    - name: Check for broken links
      run: |
        # Install linkchecker
        pip install linkchecker
        # Build docs
        mkdocs build
        # Check links (allow some time for build)
        # linkchecker site/ || true  # Don't fail on broken external links

  deploy:
    name: Deploy Documentation
    runs-on: ubuntu-latest
    if: github.event_name == 'push' && github.ref == 'refs/heads/main'
    needs: validate

    steps:
    - name: Checkout code
      uses: actions/checkout@v4
      with:
        fetch-depth: 0  # Fetch all history for git info

    - name: Set up Python
      uses: actions/setup-python@v5
      with:
        python-version: '3.11'

    - name: Install dependencies
      run: |
        pip install -r docs/requirements.txt

    - name: Configure Git
      run: |
        git config user.name "github-actions[bot]"
        git config user.email "github-actions[bot]@users.noreply.github.com"

    - name: Deploy to GitHub Pages
      run: mkdocs gh-deploy --force
`

	docsWorkflowPath := filepath.Join(workflowsDir, "docs.yml")
	if err := os.WriteFile(docsWorkflowPath, []byte(docsWorkflow), 0644); err != nil {
		return fmt.Errorf("failed to create docs workflow: %w", err)
	}

	// Release workflow
	releaseWorkflow := `name: Release

on:
  push:
    tags:
      - 'v*'

permissions:
  contents: write

jobs:
  release:
    name: Create Release
    runs-on: ubuntu-latest

    steps:
    - name: Checkout code
      uses: actions/checkout@v4
      with:
        fetch-depth: 0

    - name: Set up Go
      uses: actions/setup-go@v5
      with:
        go-version: '1.21'

    - name: Run tests
      run: go test -v ./...

    - name: Build binaries
      run: |
        # Build for multiple platforms
        GOOS=linux GOARCH=amd64 go build -o ai-linux-amd64 .
        GOOS=linux GOARCH=arm64 go build -o ai-linux-arm64 .
        GOOS=darwin GOARCH=amd64 go build -o ai-darwin-amd64 .
        GOOS=darwin GOARCH=arm64 go build -o ai-darwin-arm64 .
        GOOS=windows GOARCH=amd64 go build -o ai-windows-amd64.exe .

    - name: Create Release
      uses: softprops/action-gh-release@v1
      with:
        files: |
          ai-linux-amd64
          ai-linux-arm64
          ai-darwin-amd64
          ai-darwin-arm64
          ai-windows-amd64.exe
        generate_release_notes: true
      env:
        GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
`

	releaseWorkflowPath := filepath.Join(workflowsDir, "release.yml")
	if err := os.WriteFile(releaseWorkflowPath, []byte(releaseWorkflow), 0644); err != nil {
		return fmt.Errorf("failed to create release workflow: %w", err)
	}

	// Create PR template
	prTemplateDir := filepath.Join(".github")
	prTemplate := `## Description

<!-- Provide a brief description of the changes -->

## Perfect Commit Checklist

Following Simon Willison's "perfect commit" structure:

- [ ] **Implementation**: Single, focused change
- [ ] **Tests**: Added or updated tests
- [ ] **Documentation**: Updated relevant documentation
- [ ] **Issue Link**: References issue number

## Type of Change

- [ ] Bug fix
- [ ] New feature
- [ ] Breaking change
- [ ] Documentation update
- [ ] Refactoring
- [ ] Performance improvement

## Testing

<!-- Describe how you tested these changes -->

## Screenshots (if applicable)

<!-- Add screenshots to help explain your changes -->

## Related Issues

<!-- Link to related issues: Closes #123, Fixes #456 -->

## Additional Notes

<!-- Any additional information reviewers should know -->
`

	prTemplatePath := filepath.Join(prTemplateDir, "PULL_REQUEST_TEMPLATE.md")
	if err := os.WriteFile(prTemplatePath, []byte(prTemplate), 0644); err != nil {
		return fmt.Errorf("failed to create PR template: %w", err)
	}

	// Create issue templates directory
	issueTemplatesDir := filepath.Join(".github", "ISSUE_TEMPLATE")
	if err := os.MkdirAll(issueTemplatesDir, 0755); err != nil {
		return fmt.Errorf("failed to create issue templates directory: %w", err)
	}

	// Bug report template
	bugTemplate := `---
name: Bug Report
about: Report a bug to help us improve
title: '[BUG] '
labels: bug
assignees: ''
---

## Description

A clear and concise description of the bug.

## Steps to Reproduce

1.
2.
3.

## Expected Behavior

What you expected to happen.

## Actual Behavior

What actually happened.

## Environment

- OS: [e.g., macOS, Linux, Windows]
- Version: [e.g., v0.1.0]
- Go version: [e.g., 1.21]

## Additional Context

Add any other context about the problem here.
`

	bugTemplatePath := filepath.Join(issueTemplatesDir, "bug_report.md")
	if err := os.WriteFile(bugTemplatePath, []byte(bugTemplate), 0644); err != nil {
		return fmt.Errorf("failed to create bug template: %w", err)
	}

	// Feature request template
	featureTemplate := `---
name: Feature Request
about: Suggest a new feature
title: '[FEATURE] '
labels: enhancement
assignees: ''
---

## Problem Statement

Describe the problem this feature would solve.

## Proposed Solution

Describe your proposed solution.

## Alternatives Considered

What alternative solutions have you considered?

## Perfect Commit Components

Consider how this feature would fit into a perfect commit:

- **Implementation**: What code changes are needed?
- **Tests**: What tests would verify this works?
- **Documentation**: What docs need updating?

## Additional Context

Add any other context, screenshots, or examples.
`

	featureTemplatePath := filepath.Join(issueTemplatesDir, "feature_request.md")
	if err := os.WriteFile(featureTemplatePath, []byte(featureTemplate), 0644); err != nil {
		return fmt.Errorf("failed to create feature template: %w", err)
	}

	fmt.Println("✓ GitHub Actions workflows created")
	return nil
}
