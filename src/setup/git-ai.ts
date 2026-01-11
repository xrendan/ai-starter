import { $ } from 'bun';
import { existsSync } from 'fs';
import chalk from 'chalk';

export async function installGitAI(): Promise<void> {
  // Check if git-ai is already installed
  try {
    await $`git-ai --version`.quiet();
    console.log(chalk.green('✓ git-ai already installed'));
    return;
  } catch {
    // Not installed, proceed with installation
  }

  console.log(chalk.blue('Installing git-ai...'));

  try {
    // Detect OS
    const platform = process.platform;

    if (platform === 'win32') {
      // Windows installation
      await $`powershell -NoProfile -ExecutionPolicy Bypass -Command "irm http://usegitai.com/install.ps1 | iex"`;
    } else {
      // Unix-like systems (Linux, macOS, WSL)
      await $`curl -sSL https://usegitai.com/install.sh | bash`;
    }

    // Configure git-ai for prompt storage
    await $`git-ai config set prompt_storage notes`;

    console.log(chalk.green('✓ git-ai installed and configured'));
    console.log(chalk.gray('  Prompts will be stored in git notes'));
  } catch (error) {
    console.warn(chalk.yellow(`Warning: Failed to install git-ai: ${error}`));
    console.log(chalk.gray('  You can install it manually: https://github.com/acunniffe/git-ai'));
  }
}
