package setup

import (
	"fmt"
	"os/exec"
)

// runCommand executes a command and returns an error if it fails
func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("command failed: %w\nOutput: %s", err, output)
	}
	return nil
}
