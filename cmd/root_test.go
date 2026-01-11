package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootCommand(t *testing.T) {
	t.Run("root command executes without error", func(t *testing.T) {
		// Capture output
		buf := new(bytes.Buffer)
		rootCmd.SetOut(buf)
		rootCmd.SetErr(buf)
		rootCmd.SetArgs([]string{"--help"})

		err := rootCmd.Execute()
		assert.NoError(t, err)

		output := buf.String()
		assert.Contains(t, output, "ai is a CLI tool")
		assert.Contains(t, output, "AI-driven development")
	})

	t.Run("version flag works", func(t *testing.T) {
		buf := new(bytes.Buffer)
		rootCmd.SetOut(buf)
		rootCmd.SetErr(buf)
		rootCmd.SetArgs([]string{"--version"})

		err := rootCmd.Execute()
		assert.NoError(t, err)

		output := buf.String()
		assert.Contains(t, output, "version")
	})
}
