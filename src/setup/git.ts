import { existsSync } from 'fs';
import { $ } from 'bun';
import chalk from 'chalk';

export async function initGit(): Promise<void> {
  if (existsSync('.git')) {
    console.log(chalk.green('✓ Git repository already initialized'));
    return;
  }

  try {
    await $`git init`;
    console.log(chalk.green('✓ Git repository initialized'));
  } catch (error) {
    throw new Error(`Failed to initialize git repository: ${error}`);
  }
}
