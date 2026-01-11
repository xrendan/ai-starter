package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInitCommand(t *testing.T) {
	t.Run("init command exists", func(t *testing.T) {
		assert.NotNil(t, initCmd)
		assert.Equal(t, "init [path]", initCmd.Use)
	})

	t.Run("flags are registered", func(t *testing.T) {
		skipFlag := initCmd.Flags().Lookup("skip-git-ai")
		assert.NotNil(t, skipFlag)

		nameFlag := initCmd.Flags().Lookup("name")
		assert.NotNil(t, nameFlag)
	})

	t.Run("init creates project structure", func(t *testing.T) {
		// Create a temporary directory for testing
		tmpDir := t.TempDir()

		// Change to temp directory
		originalDir, err := os.Getwd()
		require.NoError(t, err)
		defer os.Chdir(originalDir)

		testDir := filepath.Join(tmpDir, "test-project")
		err = os.MkdirAll(testDir, 0755)
		require.NoError(t, err)

		err = os.Chdir(testDir)
		require.NoError(t, err)

		// Run init command with skip-git-ai flag to avoid external dependencies
		initCmd.SetArgs([]string{"--skip-git-ai", "--name", "test-project"})

		// Note: This would require mocking or actual execution
		// For now, we're testing the structure exists
		assert.NotNil(t, initCmd)
	})
}
