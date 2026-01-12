import { describe, test, expect, beforeEach, afterEach } from 'bun:test';
import { mkdir, rm, readFile, access } from 'node:fs/promises';
import { setupDocs } from '../src/setup/docs';
import { setupCommitTemplates } from '../src/setup/commits';
import { setupGithubActions } from '../src/setup/github';
import { setupProjectStructure } from '../src/setup/project';
import { createReadme } from '../src/setup/readme';
import { initGit } from '../src/setup/git';

const TEST_DIR = '/tmp/ai-test-' + Date.now();

beforeEach(async () => {
  await mkdir(TEST_DIR, { recursive: true });
  process.chdir(TEST_DIR);
});

afterEach(async () => {
  try {
    await rm(TEST_DIR, { recursive: true, force: true });
  } catch {
    // Ignore cleanup errors
  }
});

describe('setupDocs', () => {
  test('creates documentation structure', async () => {
    await setupDocs('test-project');

    // Check directories exist
    await access('docs');
    await access('docs/.vitepress');
    await access('docs/adr');
    await access('docs/development');
    await access('docs/guides');
    await access('docs/prompts');

    // Check VitePress config exists
    const vitepressConfig = await readFile('docs/.vitepress/config.ts', 'utf-8');
    expect(vitepressConfig).toContain('test-project');

    const indexContent = await readFile('docs/index.md', 'utf-8');
    expect(indexContent).toContain('test-project');

    // Check ADRs
    await access('docs/adr/001-use-vitepress.md');
    await access('docs/adr/004-use-bun-typescript.md');
  });
});

describe('setupCommitTemplates', () => {
  test('creates commit templates and hooks', async () => {
    await initGit();
    await setupCommitTemplates();

    // Check .gitmessage exists
    const gitmessage = await readFile('.gitmessage', 'utf-8');
    expect(gitmessage).toContain('Perfect Commit Structure');

    // Check hooks exist
    await access('.git/hooks/commit-msg');
    await access('.git/hooks/prepare-commit-msg');
    await access('.git/hooks/pre-commit');
  });
});

describe('setupGithubActions', () => {
  test('creates GitHub workflows', async () => {
    await setupGithubActions('test-project');

    // Check workflows
    await access('.github/workflows/ci.yml');
    await access('.github/workflows/docs.yml');
    await access('.github/workflows/release.yml');

    const ciContent = await readFile('.github/workflows/ci.yml', 'utf-8');
    expect(ciContent).toContain('bun');

    // Check templates
    await access('.github/PULL_REQUEST_TEMPLATE.md');
    await access('.github/ISSUE_TEMPLATE/bug_report.md');
  });
});

describe('setupProjectStructure', () => {
  test('creates project structure', async () => {
    await setupProjectStructure('test-project');

    // Check directories
    await access('src');
    await access('tests');

    // Check files
    await access('.gitignore');
    await access('Makefile');

    const gitignore = await readFile('.gitignore', 'utf-8');
    expect(gitignore).toContain('node_modules');
  });
});

describe('createReadme', () => {
  test('creates README with project name', async () => {
    await createReadme('test-project');

    const readme = await readFile('README.md', 'utf-8');
    expect(readme).toContain('test-project');
    expect(readme).toContain('Perfect Commit');
    expect(readme).toContain('OpenTUI');
  });
});
