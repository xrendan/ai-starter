package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/xrendan/ai-starter/internal/setup"
)

var (
	skipGitAI   bool
	projectName string
)

var initCmd = &cobra.Command{
	Use:   "init [path]",
	Short: "Initialize a new AI-driven repository",
	Long: `Initialize a new repository with AI-driven development best practices.

This command sets up:
- Documentation structure with MkDocs
- git-ai for prompt tracking (the new source code)
- Commit templates following Simon Willison's "perfect commit" structure
- Git hooks for commit validation
- GitHub Actions CI/CD workflows
- Testing infrastructure`,
	Args: cobra.MaximumNArgs(1),
	RunE: runInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVar(&skipGitAI, "skip-git-ai", false, "Skip git-ai installation")
	initCmd.Flags().StringVarP(&projectName, "name", "n", "", "Project name (defaults to directory name)")
}

func runInit(cmd *cobra.Command, args []string) error {
	// Determine target directory
	targetDir := "."
	if len(args) > 0 {
		targetDir = args[0]
	}

	absPath, err := filepath.Abs(targetDir)
	if err != nil {
		return fmt.Errorf("failed to resolve path: %w", err)
	}

	// Determine project name
	if projectName == "" {
		projectName = filepath.Base(absPath)
	}

	fmt.Printf("🚀 Initializing AI-driven repository: %s\n", projectName)
	fmt.Printf("📁 Target directory: %s\n\n", absPath)

	// Create directory if it doesn't exist
	if err := os.MkdirAll(absPath, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Change to target directory
	if err := os.Chdir(absPath); err != nil {
		return fmt.Errorf("failed to change directory: %w", err)
	}

	// Initialize git if not already initialized
	if err := setup.InitGit(); err != nil {
		return fmt.Errorf("git initialization failed: %w", err)
	}

	// Set up documentation structure
	fmt.Println("📚 Setting up documentation structure...")
	if err := setup.SetupDocs(projectName); err != nil {
		return fmt.Errorf("docs setup failed: %w", err)
	}

	// Set up commit templates and hooks
	fmt.Println("📝 Setting up commit templates and hooks...")
	if err := setup.SetupCommitTemplates(); err != nil {
		return fmt.Errorf("commit templates setup failed: %w", err)
	}

	// Install git-ai
	if !skipGitAI {
		fmt.Println("🤖 Installing and configuring git-ai...")
		if err := setup.InstallGitAI(); err != nil {
			fmt.Printf("⚠️  Warning: git-ai installation failed: %v\n", err)
			fmt.Println("   You can install it manually later with: curl -sSL https://usegitai.com/install.sh | bash")
		}
	}

	// Set up GitHub Actions
	fmt.Println("⚙️  Setting up GitHub Actions CI/CD...")
	if err := setup.SetupGithubActions(projectName); err != nil {
		return fmt.Errorf("GitHub Actions setup failed: %w", err)
	}

	// Set up project structure
	fmt.Println("📦 Setting up project structure...")
	if err := setup.SetupProjectStructure(projectName); err != nil {
		return fmt.Errorf("project structure setup failed: %w", err)
	}

	// Create README
	fmt.Println("📄 Creating README...")
	if err := setup.CreateReadme(projectName); err != nil {
		return fmt.Errorf("README creation failed: %w", err)
	}

	fmt.Println("\n✅ Repository initialized successfully!")
	fmt.Println("\n📖 Next steps:")
	fmt.Println("   1. Review the generated documentation in docs/")
	fmt.Println("   2. Update docs/adr/ with your architectural decisions")
	fmt.Println("   3. Follow the commit template when making changes")
	fmt.Println("   4. Run 'mkdocs serve' to preview documentation")
	fmt.Println("   5. Push to GitHub to trigger CI/CD workflows")

	return nil
}
