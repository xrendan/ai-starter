# ai-starter

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
git clone https://github.com/xrendan/ai-starter.git
cd ai-starter
go build -o ai .
sudo mv ai /usr/local/bin/
\'\'\'

### From Release

Download the latest release for your platform from the [releases page](https://github.com/xrendan/ai-starter/releases).

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

- 📖 [Documentation](https://xrendan.github.io/ai-starter/)
- 🐛 [Issue Tracker](https://github.com/xrendan/ai-starter/issues)
- 💬 [Discussions](https://github.com/xrendan/ai-starter/discussions)
