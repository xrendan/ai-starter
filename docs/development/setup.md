# Setup Guide

## Prerequisites

- Go 1.21 or higher
- Git
- Python 3.8+ (for documentation)
- Make (optional)

## Installation

1. Clone the repository:
   \'\'\'bash
   git clone https://github.com/xrendan/ai-starter.git
   cd ai-starter
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
