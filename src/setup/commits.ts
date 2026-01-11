import { writeFile, mkdir, chmod } from 'fs/promises';
import { join } from 'path';
import { $ } from 'bun';
import chalk from 'chalk';

export async function setupCommitTemplates(): Promise<void> {
  const commitTemplate = `# [Type] Brief description (#issue)
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
`;

  await writeFile('.gitmessage', commitTemplate);

  // Configure git to use the template
  try {
    await $`git config commit.template .gitmessage`;
    await $`git config commit.cleanup strip`;
  } catch (error) {
    console.warn(chalk.yellow(`Warning: failed to configure git: ${error}`));
  }

  // Create hooks directory
  const hooksDir = join('.git', 'hooks');
  await mkdir(hooksDir, { recursive: true });

  // Create commit-msg hook
  const commitMsgHook = `#!/bin/bash
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

# Check if commit message is empty or only contains comments
if [ -z "$(echo "$COMMIT_MSG" | grep -v '^#')" ]; then
    echo "Error: Empty commit message"
    exit 1
fi

exit 0
`;

  const commitMsgPath = join(hooksDir, 'commit-msg');
  await writeFile(commitMsgPath, commitMsgHook);
  await chmod(commitMsgPath, 0o755);

  // Create prepare-commit-msg hook
  const prepareCommitMsgHook = `#!/bin/bash
#
# Prepare commit message hook
# Auto-adds issue number from branch name
#

COMMIT_MSG_FILE=$1
COMMIT_SOURCE=$2

# Only run for regular commits (not merge, squash, etc.)
if [ -z "$COMMIT_SOURCE" ] || [ "$COMMIT_SOURCE" = "message" ]; then
    BRANCH_NAME=$(git symbolic-ref --short HEAD 2>/dev/null)

    # Extract issue number from branch name (e.g., feature/123-description -> #123)
    ISSUE_NUM=$(echo "$BRANCH_NAME" | grep -oE '[0-9]+' | head -1)

    if [ -n "$ISSUE_NUM" ]; then
        # Check if issue number already in commit message
        if ! grep -q "#$ISSUE_NUM" "$COMMIT_MSG_FILE"; then
            # Add issue number to first line if not present
            sed -i "1s/$/ (#$ISSUE_NUM)/" "$COMMIT_MSG_FILE"
        fi
    fi
fi
`;

  const prepareCommitMsgPath = join(hooksDir, 'prepare-commit-msg');
  await writeFile(prepareCommitMsgPath, prepareCommitMsgHook);
  await chmod(prepareCommitMsgPath, 0o755);

  // Create pre-commit hook
  const preCommitHook = `#!/bin/bash
#
# Pre-commit hook
# Reminds about perfect commit structure
#

echo "📋 Perfect Commit Checklist:"
echo "   ✓ Implementation: Single, focused change?"
echo "   ✓ Tests: Added or updated tests?"
echo "   ✓ Documentation: Updated relevant docs?"
echo "   ✓ Issue: Linked to issue number?"

exit 0
`;

  const preCommitPath = join(hooksDir, 'pre-commit');
  await writeFile(preCommitPath, preCommitHook);
  await chmod(preCommitPath, 0o755);

  console.log(chalk.green('✓ Commit templates and hooks configured'));
}
