package setup

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// InitGit initializes a git repository if not already initialized
func InitGit() error {
	// Check if .git exists
	if _, err := os.Stat(".git"); err == nil {
		fmt.Println("✓ Git repository already initialized")
		return nil
	}

	cmd := exec.Command("git", "init")
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git init failed: %w\nOutput: %s", err, output)
	}

	fmt.Println("✓ Git repository initialized")
	return nil
}

// InstallGitAI installs and configures git-ai for prompt tracking
func InstallGitAI() error {
	// Check if git-ai is already installed
	if _, err := exec.LookPath("git-ai"); err == nil {
		fmt.Println("✓ git-ai already installed")
		return configureGitAI()
	}

	// Install git-ai based on OS
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		// Use PowerShell for Windows
		script := `irm http://usegitai.com/install.ps1 | iex`
		cmd = exec.Command("powershell", "-NoProfile", "-ExecutionPolicy", "Bypass", "-Command", script)
	default:
		// Unix-like systems (Linux, macOS, WSL)
		script := `curl -sSL https://usegitai.com/install.sh | bash`
		cmd = exec.Command("bash", "-c", script)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("installation failed: %w\nOutput: %s", err, output)
	}

	fmt.Println("✓ git-ai installed successfully")
	return configureGitAI()
}

// configureGitAI sets up git-ai configuration
func configureGitAI() error {
	// Install hooks
	cmd := exec.Command("git-ai", "install-hooks")
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("hook installation failed: %w\nOutput: %s", err, output)
	}

	// Configure prompt storage in git notes
	cmd = exec.Command("git-ai", "config", "set", "prompt_storage", "notes")
	if output, err := cmd.CombinedOutput(); err != nil {
		// This might fail if the command syntax is different, log but don't error
		fmt.Printf("Note: prompt_storage config may need manual setup: %s\n", output)
	}

	fmt.Println("✓ git-ai configured for prompt tracking")
	return nil
}
