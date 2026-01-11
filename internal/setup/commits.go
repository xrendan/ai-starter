package setup

import (
	"fmt"
	"os"
	"path/filepath"
)

// SetupCommitTemplates creates commit message template and hooks
func SetupCommitTemplates() error {
	// Create commit message template following Simon Willison's perfect commit structure
	commitTemplate := `# [Type] Brief description (#issue)
#
# Type: feat|fix|docs|refactor|test|chore
#
# Perfect Commit Structure (Simon Willison):
# 1. Implementation - a single, focused change
# 2. Tests - demonstrate it works
# 3. Documentation - keep docs in sync
# 4. Issue Link - provide context
#
# ============================================

# Implementation:
# - What changed and why
# -

# Tests:
# - What tests were added/modified
# -

# Documentation:
# - What docs were updated
# -

# Issue Context:
# - Reference issue number
# - Closes #

# ============================================
# Example:
# feat: Add user authentication (#42)
#
# Implementation:
# - Added JWT-based authentication middleware
# - Implemented login and logout endpoints
#
# Tests:
# - Added auth middleware tests
# - Added integration tests for login flow
#
# Documentation:
# - Updated API documentation with auth endpoints
# - Added authentication guide
#
# Closes #42
# ============================================
`

	if err := os.WriteFile(".gitmessage", []byte(commitTemplate), 0644); err != nil {
		return fmt.Errorf("failed to create .gitmessage: %w", err)
	}

	// Configure git to use the template
	gitConfigCmd := [][]string{
		{"git", "config", "commit.template", ".gitmessage"},
		{"git", "config", "commit.cleanup", "strip"},
	}

	for _, args := range gitConfigCmd {
		if err := runCommand(args[0], args[1:]...); err != nil {
			fmt.Printf("Warning: failed to configure git: %v\n", err)
		}
	}

	// Create hooks directory
	hooksDir := filepath.Join(".git", "hooks")
	if err := os.MkdirAll(hooksDir, 0755); err != nil {
		return fmt.Errorf("failed to create hooks directory: %w", err)
	}

	// Create commit-msg hook
	commitMsgHook := `#!/bin/bash
#
# Commit message validation hook
# Validates commit message structure for "perfect commits"
#

COMMIT_MSG_FILE=$1
COMMIT_MSG=$(cat "$COMMIT_MSG_FILE")

# Skip validation for merge commits
if grep -q "^Merge" "$COMMIT_MSG_FILE"; then
    exit 0
fi

# Skip validation for revert commits
if grep -q "^Revert" "$COMMIT_MSG_FILE"; then
    exit 0
fi

# Check for commit message structure
# Allow simple commits but encourage perfect structure

# Warn if commit doesn't follow recommended structure
if ! echo "$COMMIT_MSG" | grep -q "Implementation:\|Tests:\|Documentation:"; then
    echo "⚠️  Reminder: Consider using the perfect commit structure:"
    echo "   - Implementation: what changed"
    echo "   - Tests: what tests were added"
    echo "   - Documentation: what docs were updated"
    echo "   - Issue reference: #number"
    echo ""
    echo "   This is a recommendation, not a requirement."
    echo "   Simple commits (like typo fixes) can skip this structure."
    echo ""
fi

# Check for overly short commit messages (< 10 chars)
MSG_LENGTH=$(echo "$COMMIT_MSG" | head -1 | wc -c)
if [ "$MSG_LENGTH" -lt 10 ]; then
    echo "❌ Error: Commit message too short (< 10 characters)"
    echo "   Please provide a meaningful commit message."
    exit 1
fi

exit 0
`

	commitMsgHookPath := filepath.Join(hooksDir, "commit-msg")
	if err := os.WriteFile(commitMsgHookPath, []byte(commitMsgHook), 0755); err != nil {
		return fmt.Errorf("failed to create commit-msg hook: %w", err)
	}

	// Create prepare-commit-msg hook for issue reference
	prepareCommitMsgHook := `#!/bin/bash
#
# Prepare commit message hook
# Automatically adds branch issue number if in format: feature/123-description
#

COMMIT_MSG_FILE=$1
COMMIT_SOURCE=$2

# Only add issue number for regular commits (not merge, squash, etc.)
if [ -z "$COMMIT_SOURCE" ]; then
    BRANCH_NAME=$(git symbolic-ref --short HEAD 2>/dev/null)

    # Extract issue number from branch name (e.g., feature/123-description -> #123)
    ISSUE_NUMBER=$(echo "$BRANCH_NAME" | sed -n 's/.*\/\([0-9]\+\).*/\1/p')

    if [ -n "$ISSUE_NUMBER" ]; then
        # Check if issue number is already in the commit message
        if ! grep -q "#$ISSUE_NUMBER" "$COMMIT_MSG_FILE"; then
            # Add issue number to first line if not present
            sed -i.bak "1s/$/ (#$ISSUE_NUMBER)/" "$COMMIT_MSG_FILE"
            rm -f "${COMMIT_MSG_FILE}.bak"
        fi
    fi
fi
`

	prepareCommitMsgHookPath := filepath.Join(hooksDir, "prepare-commit-msg")
	if err := os.WriteFile(prepareCommitMsgHookPath, []byte(prepareCommitMsgHook), 0755); err != nil {
		return fmt.Errorf("failed to create prepare-commit-msg hook: %w", err)
	}

	// Create pre-commit hook reminder
	preCommitHook := `#!/bin/bash
#
# Pre-commit hook
# Reminds about perfect commit checklist
#

echo "📋 Perfect Commit Checklist:"
echo "   ✓ Implementation: Single, focused change?"
echo "   ✓ Tests: Added or updated tests?"
echo "   ✓ Documentation: Updated relevant docs?"
echo "   ✓ Issue: Linked to issue number?"
echo ""
`

	preCommitHookPath := filepath.Join(hooksDir, "pre-commit")
	if err := os.WriteFile(preCommitHookPath, []byte(preCommitHook), 0755); err != nil {
		return fmt.Errorf("failed to create pre-commit hook: %w", err)
	}

	fmt.Println("✓ Commit templates and hooks configured")
	return nil
}
