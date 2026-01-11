package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ai",
	Short: "AI-driven repository setup and management tool",
	Long: `ai is a CLI tool designed to set up and manage repositories
optimized for AI-driven development workflows.

It enforces best practices including:
- Documentation-driven development
- Simon Willison's "perfect commit" structure
- AI prompt tracking with git-ai
- Automated testing and CI/CD`,
	Version: "0.1.0",
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.SetOut(os.Stdout)
	rootCmd.SetErr(os.Stderr)
}
