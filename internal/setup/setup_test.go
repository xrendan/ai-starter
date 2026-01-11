package setup

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSetupDocs(t *testing.T) {
	tmpDir := t.TempDir()
	originalDir, err := os.Getwd()
	require.NoError(t, err)
	defer os.Chdir(originalDir)

	err = os.Chdir(tmpDir)
	require.NoError(t, err)

	err = SetupDocs("test-project")
	require.NoError(t, err)

	// Check that directories were created
	dirs := []string{
		"docs",
		"docs/adr",
		"docs/api",
		"docs/guides",
		"docs/prompts",
		"docs/development",
	}

	for _, dir := range dirs {
		assert.DirExists(t, dir, "Directory %s should exist", dir)
	}

	// Check that key files were created
	files := []string{
		"mkdocs.yml",
		"docs/index.md",
		"docs/adr/index.md",
		"docs/adr/001-use-mkdocs.md",
		"docs/adr/002-perfect-commits.md",
		"docs/adr/003-track-ai-prompts.md",
		"docs/development/setup.md",
		"docs/development/contributing.md",
		"docs/development/testing.md",
		"docs/prompts/index.md",
		"docs/guides/index.md",
		"docs/requirements.txt",
	}

	for _, file := range files {
		assert.FileExists(t, file, "File %s should exist", file)
	}

	// Verify mkdocs.yml contains project name
	content, err := os.ReadFile("mkdocs.yml")
	require.NoError(t, err)
	assert.Contains(t, string(content), "test-project")
}

func TestSetupCommitTemplates(t *testing.T) {
	tmpDir := t.TempDir()
	originalDir, err := os.Getwd()
	require.NoError(t, err)
	defer os.Chdir(originalDir)

	err = os.Chdir(tmpDir)
	require.NoError(t, err)

	// Initialize git
	err = InitGit()
	require.NoError(t, err)

	err = SetupCommitTemplates()
	require.NoError(t, err)

	// Check that .gitmessage was created
	assert.FileExists(t, ".gitmessage")

	content, err := os.ReadFile(".gitmessage")
	require.NoError(t, err)
	assert.Contains(t, string(content), "Perfect Commit Structure")
	assert.Contains(t, string(content), "Implementation:")
	assert.Contains(t, string(content), "Tests:")
	assert.Contains(t, string(content), "Documentation:")

	// Check that hooks were created
	hooks := []string{
		".git/hooks/commit-msg",
		".git/hooks/prepare-commit-msg",
		".git/hooks/pre-commit",
	}

	for _, hook := range hooks {
		assert.FileExists(t, hook, "Hook %s should exist", hook)

		// Check that hooks are executable
		info, err := os.Stat(hook)
		require.NoError(t, err)
		assert.NotEqual(t, 0, info.Mode()&0111, "Hook %s should be executable", hook)
	}
}

func TestSetupGithubActions(t *testing.T) {
	tmpDir := t.TempDir()
	originalDir, err := os.Getwd()
	require.NoError(t, err)
	defer os.Chdir(originalDir)

	err = os.Chdir(tmpDir)
	require.NoError(t, err)

	err = SetupGithubActions("test-project")
	require.NoError(t, err)

	// Check that workflows were created
	workflows := []string{
		".github/workflows/ci.yml",
		".github/workflows/docs.yml",
		".github/workflows/release.yml",
	}

	for _, workflow := range workflows {
		assert.FileExists(t, workflow, "Workflow %s should exist", workflow)
	}

	// Check that templates were created
	assert.FileExists(t, ".github/PULL_REQUEST_TEMPLATE.md")
	assert.FileExists(t, ".github/ISSUE_TEMPLATE/bug_report.md")
	assert.FileExists(t, ".github/ISSUE_TEMPLATE/feature_request.md")

	// Verify CI workflow content
	ciContent, err := os.ReadFile(".github/workflows/ci.yml")
	require.NoError(t, err)
	assert.Contains(t, string(ciContent), "go test")
	assert.Contains(t, string(ciContent), "golangci-lint")
}

func TestSetupProjectStructure(t *testing.T) {
	tmpDir := t.TempDir()
	originalDir, err := os.Getwd()
	require.NoError(t, err)
	defer os.Chdir(originalDir)

	err = os.Chdir(tmpDir)
	require.NoError(t, err)

	err = SetupProjectStructure("test-project")
	require.NoError(t, err)

	// Check that directories were created
	dirs := []string{
		"cmd",
		"internal",
		"pkg",
		"test",
		"scripts",
	}

	for _, dir := range dirs {
		assert.DirExists(t, dir, "Directory %s should exist", dir)
	}

	// Check that files were created
	files := []string{
		".gitignore",
		"Makefile",
		".golangci.yml",
		"CONTRIBUTING.md",
	}

	for _, file := range files {
		assert.FileExists(t, file, "File %s should exist", file)
	}

	// Verify .gitignore content
	gitignoreContent, err := os.ReadFile(".gitignore")
	require.NoError(t, err)
	assert.Contains(t, string(gitignoreContent), "*.exe")
	assert.Contains(t, string(gitignoreContent), ".DS_Store")
}

func TestCreateReadme(t *testing.T) {
	tmpDir := t.TempDir()
	originalDir, err := os.Getwd()
	require.NoError(t, err)
	defer os.Chdir(originalDir)

	err = os.Chdir(tmpDir)
	require.NoError(t, err)

	err = CreateReadme("test-project")
	require.NoError(t, err)

	assert.FileExists(t, "README.md")

	content, err := os.ReadFile("README.md")
	require.NoError(t, err)
	assert.Contains(t, string(content), "test-project")
	assert.Contains(t, string(content), "Perfect Commits")
	assert.Contains(t, string(content), "Simon Willison")
}

func TestInitGit(t *testing.T) {
	tmpDir := t.TempDir()
	originalDir, err := os.Getwd()
	require.NoError(t, err)
	defer os.Chdir(originalDir)

	err = os.Chdir(tmpDir)
	require.NoError(t, err)

	err = InitGit()
	require.NoError(t, err)

	// Check that .git directory was created
	assert.DirExists(t, ".git")

	// Running again should not error
	err = InitGit()
	require.NoError(t, err)
}
