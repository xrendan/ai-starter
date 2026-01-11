import { writeFile } from 'node:fs/promises';
import chalk from 'chalk';

export async function createReadme(projectName: string): Promise<void> {
  const readme = `# ${projectName}

AI-driven repository setup and management tool built with Bun, TypeScript, and OpenTUI.

## Features

- 🚀 **Fast Setup**: Initialize repositories with best practices in seconds
- 📚 **Documentation-First**: MkDocs with Material theme
- ✨ **Perfect Commits**: Following Simon Willison's commit structure
- 🤖 **AI Prompt Tracking**: git-ai integration for tracking AI-generated code
- 🧪 **Testing**: Built-in Bun test framework
- 🔄 **CI/CD**: GitHub Actions workflows included
- 🎨 **Interactive TUI**: Built with OpenTUI for rich terminal interfaces

## Installation

### Prerequisites

- [Bun](https://bun.sh/) >= 1.0.0
- [Zig](https://ziglang.org/) >= 0.13.0 (for OpenTUI)
- Git
- Python 3.8+ (for MkDocs documentation)

### Install from source

\`\`\`bash
git clone https://github.com/xrendan/${projectName}.git
cd ${projectName}
bun install
bun run build:binary
\`\`\`

### Install globally

\`\`\`bash
bun install -g @xrendan/ai-starter
\`\`\`

## Usage

### Initialize a new repository

\`\`\`bash
ai init [path]
\`\`\`

### Options

- \`-n, --name <name>\` - Project name (defaults to directory name)
- \`--skip-git-ai\` - Skip git-ai installation

### Examples

\`\`\`bash
# Initialize in current directory
ai init

# Initialize in a specific directory
ai init my-project

# Initialize with custom name
ai init my-project --name "My Awesome Project"

# Initialize without git-ai
ai init --skip-git-ai
\`\`\`

## What Gets Created

When you run \`ai init\`, the following structure is created:

### Documentation
- \`docs/\` - Complete MkDocs documentation structure
  - \`adr/\` - Architectural Decision Records
  - \`development/\` - Setup, contributing, and testing guides
  - \`guides/\` - User guides
  - \`prompts/\` - AI prompts tracking
- \`mkdocs.yml\` - MkDocs configuration with Material theme

### Git Configuration
- \`.gitmessage\` - Commit message template
- \`.git/hooks/\` - Git hooks for commit validation
  - \`commit-msg\` - Validates commit messages
  - \`prepare-commit-msg\` - Auto-adds issue numbers
  - \`pre-commit\` - Shows perfect commit checklist

### GitHub
- \`.github/workflows/\` - CI/CD workflows
  - \`ci.yml\` - Testing and linting
  - \`docs.yml\` - Documentation deployment
  - \`release.yml\` - Multi-platform binary builds
- \`.github/PULL_REQUEST_TEMPLATE.md\` - PR template
- \`.github/ISSUE_TEMPLATE/\` - Issue templates

### Project Files
- \`src/\` - Source code directory
- \`tests/\` - Test directory
- \`package.json\` - Project configuration
- \`tsconfig.json\` - TypeScript configuration
- \`.gitignore\` - Git ignore rules
- \`Makefile\` - Common development tasks
- \`.eslintrc.json\` - ESLint configuration
- \`.prettierrc\` - Prettier configuration

## Perfect Commit Structure

This project follows Simon Willison's "perfect commit" philosophy. Each commit should include:

1. **Implementation**: A single, focused change
2. **Tests**: Demonstrate the change works
3. **Documentation**: Keep docs in sync with code
4. **Issue Link**: Provide context with issue reference

Example:

\`\`\`
feat: Add user authentication (#42)

Implementation:
- Added JWT-based authentication middleware
- Implemented login and logout endpoints

Tests:
- Added auth middleware tests
- Added integration tests for login flow

Documentation:
- Updated API documentation with auth endpoints
- Added authentication guide

Closes #42
\`\`\`

## Development

\`\`\`bash
# Install dependencies
bun install

# Run in development mode
bun run dev

# Run tests
bun test

# Run tests in watch mode
bun test --watch

# Build
bun run build

# Build binary
bun run build:binary

# Lint
bun run lint

# Format code
bun run format
\`\`\`

## Documentation

Preview documentation locally:

\`\`\`bash
mkdocs serve
\`\`\`

Then visit http://localhost:8000

## Technologies

- **Runtime**: [Bun](https://bun.sh/) - Fast JavaScript runtime with built-in TypeScript support
- **UI**: [OpenTUI](https://github.com/sst/opentui) - Terminal user interface library
- **Documentation**: [MkDocs](https://www.mkdocs.org/) with [Material theme](https://squidfunk.github.io/mkdocs-material/)
- **AI Tracking**: [git-ai](https://github.com/acunniffe/git-ai) - Track AI-generated code
- **Testing**: [Bun Test](https://bun.sh/docs/cli/test) - Built-in test runner

## License

MIT

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for contribution guidelines.

## Related Projects

- [git-ai](https://github.com/acunniffe/git-ai) - Track AI code in repositories
- [Simon Willison's Perfect Commit](https://simonwillison.net/2022/Oct/29/the-perfect-commit/) - Commit structure philosophy
- [OpenTUI](https://github.com/sst/opentui) - Terminal UI framework

## Acknowledgments

- Simon Willison for the perfect commit philosophy
- The Bun team for the amazing runtime
- The OpenTUI team for the TUI framework
- The git-ai project for AI code tracking
`;

  await writeFile('README.md', readme);
  console.log(chalk.green('✓ README created'));
}
