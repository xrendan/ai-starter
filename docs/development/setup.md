# Setup Guide

## Prerequisites

- Bun 1.0 or higher
- Git
- Make (optional)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/xrendan/test-project.git
   cd test-project
   ```

2. Install dependencies:
   ```bash
   bun install
   ```

3. Install git-ai for prompt tracking:
   ```bash
   curl -sSL https://usegitai.com/install.sh | bash
   ```

## Development

### Building

```bash
bun run build
```

### Building Binary

```bash
bun run build:binary
```

### Testing

```bash
bun test
```

### Documentation

Preview documentation locally:

```bash
bun run docs:dev
```

The dev server will start and automatically open in your browser.

## Commit Guidelines

This project follows the "perfect commit" structure:

1. Each commit should be a focused, atomic change
2. Include tests for new functionality
3. Update documentation in the same commit
4. Reference the issue number in the commit message

Use the commit template:

```bash
git config commit.template .gitmessage
```
