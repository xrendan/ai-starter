import { writeFile, mkdir } from 'fs/promises';
import chalk from 'chalk';

export async function setupProjectStructure(projectName: string): Promise<void> {
  // Create directory structure
  const dirs = ['src', 'src/lib', 'tests', 'scripts'];

  for (const dir of dirs) {
    await mkdir(dir, { recursive: true });
  }

  // Create .gitignore
  const gitignore = `# Dependencies
node_modules/
bun.lockb

# Build outputs
dist/
*.exe
*.dll
*.so
*.dylib
ai
ai-*

# Test coverage
coverage/
*.lcov

# Environment variables
.env
.env.local
.env.*.local

# IDE
.vscode/
.idea/
*.swp
*.swo
*~
.DS_Store

# Logs
logs/
*.log
npm-debug.log*
yarn-debug.log*
yarn-error.log*

# Temporary files
*.tmp
tmp/
temp/

# Documentation build
site/
`;

  await writeFile('.gitignore', gitignore);

  // Create Makefile
  const makefile = `.PHONY: help install build test lint format clean

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \\033[36m%-15s\\033[0m %s\\n", $$1, $$2}' $(MAKEFILE_LIST)

install: ## Install dependencies
	bun install

build: ## Build the project
	bun run build

build-binary: ## Build standalone binary
	bun run build:binary

test: ## Run tests
	bun test

test-watch: ## Run tests in watch mode
	bun test --watch

lint: ## Run linter
	bun run lint

format: ## Format code
	bun run format

clean: ## Clean build artifacts
	bun run clean

dev: ## Run in development mode
	bun run dev

docs: ## Serve documentation locally
	mkdocs serve

docs-build: ## Build documentation
	mkdocs build
`;

  await writeFile('Makefile', makefile);

  // Create .eslintrc.json
  const eslintConfig = `{
  "parser": "@typescript-eslint/parser",
  "extends": [
    "eslint:recommended",
    "plugin:@typescript-eslint/recommended"
  ],
  "plugins": ["@typescript-eslint"],
  "env": {
    "node": true,
    "es2022": true
  },
  "parserOptions": {
    "ecmaVersion": 2022,
    "sourceType": "module",
    "project": "./tsconfig.json"
  },
  "rules": {
    "@typescript-eslint/no-unused-vars": "error",
    "@typescript-eslint/no-explicit-any": "warn",
    "@typescript-eslint/explicit-function-return-type": "off"
  }
}
`;

  await writeFile('.eslintrc.json', eslintConfig);

  // Create .prettierrc
  const prettierConfig = `{
  "semi": true,
  "trailingComma": "es5",
  "singleQuote": true,
  "printWidth": 100,
  "tabWidth": 2
}
`;

  await writeFile('.prettierrc', prettierConfig);

  // Create CONTRIBUTING.md symlink
  const contributing = `# Contributing

Please see [docs/development/contributing.md](docs/development/contributing.md) for contribution guidelines.
`;

  await writeFile('CONTRIBUTING.md', contributing);

  // Create example test
  const exampleTest = `import { describe, test, expect } from 'bun:test';

describe('example', () => {
  test('placeholder test', () => {
    expect(true).toBe(true);
  });
});
`;

  await writeFile('tests/example.test.ts', exampleTest);

  console.log(chalk.green('✓ Project structure created'));
}
