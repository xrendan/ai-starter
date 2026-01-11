import { Command } from 'commander';
import { resolve, basename } from 'path';
import { initGit } from '../setup/git';
import { setupDocs } from '../setup/docs';
import { setupCommitTemplates } from '../setup/commits';
import { setupGithubActions } from '../setup/github';
import { setupProjectStructure } from '../setup/project';
import { createReadme } from '../setup/readme';
import { installGitAI } from '../setup/git-ai';
import chalk from 'chalk';

interface InitOptions {
  name?: string;
  skipGitAi?: boolean;
}

export const initCommand = new Command('init')
  .description('Initialize a new repository with AI-driven development best practices')
  .argument('[path]', 'Path to initialize (defaults to current directory)', '.')
  .option('-n, --name <name>', 'Project name (defaults to directory name)')
  .option('--skip-git-ai', 'Skip git-ai installation', false)
  .action(async (path: string, options: InitOptions) => {
    try {
      const absPath = resolve(path);
      const projectName = options.name || basename(absPath);

      console.log(chalk.cyan(`🚀 Initializing AI-driven repository: ${projectName}`));
      console.log(chalk.gray(`📁 Target directory: ${absPath}\n`));

      // Change to target directory
      process.chdir(absPath);

      // Initialize git
      await initGit();

      // Setup documentation
      console.log(chalk.blue('📚 Setting up documentation structure...'));
      await setupDocs(projectName);

      // Setup commit templates and hooks
      console.log(chalk.blue('📝 Setting up commit templates and hooks...'));
      await setupCommitTemplates();

      // Setup GitHub Actions
      console.log(chalk.blue('⚙️  Setting up GitHub Actions CI/CD...'));
      await setupGithubActions(projectName);

      // Setup project structure
      console.log(chalk.blue('📦 Setting up project structure...'));
      await setupProjectStructure(projectName);

      // Create README
      console.log(chalk.blue('📄 Creating README...'));
      await createReadme(projectName);

      // Install git-ai if not skipped
      if (!options.skipGitAi) {
        console.log(chalk.blue('🤖 Installing git-ai for prompt tracking...'));
        await installGitAI();
      }

      console.log(chalk.green('\n✅ Repository initialized successfully!\n'));
      console.log(chalk.yellow('📖 Next steps:'));
      console.log('   1. Review the generated documentation in docs/');
      console.log('   2. Update docs/adr/ with your architectural decisions');
      console.log('   3. Follow the commit template when making changes');
      console.log('   4. Run \'mkdocs serve\' to preview documentation');
      if (!options.skipGitAi) {
        console.log('   5. git-ai is configured to track AI-generated code');
      }
      console.log('   6. Push to GitHub to trigger CI/CD workflows');
    } catch (error) {
      console.error(chalk.red(`\n❌ Error: ${error instanceof Error ? error.message : String(error)}`));
      process.exit(1);
    }
  });
