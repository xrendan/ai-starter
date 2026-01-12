import { mkdir, writeFile } from 'node:fs/promises';
import { join } from 'node:path';
import chalk from 'chalk';

export async function setupGithubActions(_projectName: string): Promise<void> {
  const workflowsDir = join('.github', 'workflows');
  await mkdir(workflowsDir, { recursive: true });

  // Create CI workflow
  const ciWorkflow = `name: CI

on:
  push:
    branches: [ main, develop ]
  pull_request:
    branches: [ main ]

jobs:
  test:
    runs-on: ubuntu-latest

    steps:
    - uses: actions/checkout@v4

    - name: Setup Bun
      uses: oven-sh/setup-bun@v1
      with:
        bun-version: latest

    - name: Setup Zig
      uses: goto-bus-stop/setup-zig@v2
      with:
        version: 0.13.0

    - name: Install dependencies
      run: bun install

    - name: Run tests
      run: bun test

    - name: Build
      run: bun run build

    - name: Lint
      run: bun run lint
`;

  await writeFile(join(workflowsDir, 'ci.yml'), ciWorkflow);

  // Create docs workflow
  const docsWorkflow = `name: Documentation

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

permissions:
  contents: read
  pages: write
  id-token: write

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Setup Python
        uses: actions/setup-python@v5
        with:
          python-version: '3.x'

      - name: Install dependencies
        run: |
          pip install -r docs/requirements.txt

      - name: Build documentation
        run: mkdocs build --strict

      - name: Upload artifact
        uses: actions/upload-pages-artifact@v3
        with:
          path: site

  deploy:
    environment:
      name: github-pages
      url: \${{ steps.deployment.outputs.page_url }}
    runs-on: ubuntu-latest
    needs: build
    if: github.ref == 'refs/heads/main'
    steps:
      - name: Deploy to GitHub Pages
        id: deployment
        uses: actions/deploy-pages@v4
`;

  await writeFile(join(workflowsDir, 'docs.yml'), docsWorkflow);

  // Create release workflow
  const releaseWorkflow = `name: Release

on:
  push:
    tags:
      - 'v*'

jobs:
  build:
    name: Build binaries
    runs-on: \${{ matrix.os }}
    strategy:
      matrix:
        include:
          - os: ubuntu-latest
            target: linux-x64
          - os: macos-latest
            target: darwin-x64
          - os: windows-latest
            target: win-x64

    steps:
    - uses: actions/checkout@v4

    - name: Setup Bun
      uses: oven-sh/setup-bun@v1
      with:
        bun-version: latest

    - name: Setup Zig
      uses: goto-bus-stop/setup-zig@v2
      with:
        version: 0.13.0

    - name: Install dependencies
      run: bun install

    - name: Build binary
      run: bun run build:binary

    - name: Upload artifacts
      uses: actions/upload-artifact@v4
      with:
        name: ai-\${{ matrix.target }}
        path: ai*

  release:
    needs: build
    runs-on: ubuntu-latest
    steps:
    - name: Download artifacts
      uses: actions/download-artifact@v4

    - name: Create Release
      uses: softprops/action-gh-release@v1
      with:
        files: |
          ai-*/ai*
`;

  await writeFile(join(workflowsDir, 'release.yml'), releaseWorkflow);

  // Create PR template
  const prTemplate = `## Description

Please include a summary of the change and which issue is fixed.

Fixes # (issue)

## Type of change

- [ ] Bug fix
- [ ] New feature  
- [ ] Breaking change
- [ ] Documentation update

## Perfect Commit Checklist

- [ ] Implementation: Single, focused change
- [ ] Tests: Added or updated tests
- [ ] Documentation: Updated relevant docs
- [ ] Issue: Linked to issue number

## How Has This Been Tested?

Please describe the tests that you ran to verify your changes.

- [ ] Test A
- [ ] Test B

## Additional Notes

Add any additional notes or context here.
`;

  await mkdir('.github', { recursive: true });
  await writeFile(join('.github', 'PULL_REQUEST_TEMPLATE.md'), prTemplate);

  // Create issue templates
  const bugTemplate = `---
name: Bug Report
about: Create a report to help us improve
title: '[BUG] '
labels: bug
assignees: ''
---

**Describe the bug**
A clear and concise description of what the bug is.

**To Reproduce**
Steps to reproduce the behavior:
1. Go to '...'
2. Click on '....'
3. See error

**Expected behavior**
A clear and concise description of what you expected to happen.

**Environment:**
 - OS: [e.g. Ubuntu 22.04]
 - Bun version: [e.g. 1.0.0]
 - Tool version: [e.g. 0.1.0]

**Additional context**
Add any other context about the problem here.
`;

  const featureTemplate = `---
name: Feature Request
about: Suggest an idea for this project
title: '[FEATURE] '
labels: enhancement
assignees: ''
---

**Is your feature request related to a problem?**
A clear and concise description of what the problem is.

**Describe the solution you'd like**
A clear and concise description of what you want to happen.

**Describe alternatives you've considered**
A clear and concise description of any alternative solutions or features you've considered.

**Additional context**
Add any other context or screenshots about the feature request here.
`;

  const issueTemplatesDir = join('.github', 'ISSUE_TEMPLATE');
  await mkdir(issueTemplatesDir, { recursive: true });
  await writeFile(join(issueTemplatesDir, 'bug_report.md'), bugTemplate);
  await writeFile(join(issueTemplatesDir, 'feature_request.md'), featureTemplate);

  console.log(chalk.green('✓ GitHub Actions workflows created'));
}
