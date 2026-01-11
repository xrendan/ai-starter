# Contributing Guide

Thank you for contributing! This guide will help you make effective contributions.

## Perfect Commit Structure

Following Simon Willison's philosophy, each commit should include:

### 1. Implementation
- A single, focused change
- Should be atomic and deployable
- Easy to review and understand

### 2. Tests
- Demonstrate the change works
- Prevent regressions
- Document expected behavior

### 3. Documentation
- Update relevant docs in the same commit
- Include API docs, user guides, or ADRs as needed
- Keep docs in sync with code

### 4. Issue Link
- Reference the issue number
- Provides context and discussion history
- Explains the "why" behind the change

## Workflow

1. **Create an issue** describing the change
2. **Create a branch** from main
3. **Make your changes** following the commit structure
4. **Run tests** to ensure everything works
5. **Update documentation** in the same commit
6. **Submit a PR** referencing the issue
7. **Address review feedback**

## Commit Message Format

\'\'\'
[Type] Brief description (#issue)

Implementation:
- What changed

Tests:
- What tests were added/modified

Documentation:
- What docs were updated

Closes #issue
\'\'\'

## Types

- **feat**: New feature
- **fix**: Bug fix
- **docs**: Documentation only
- **refactor**: Code refactoring
- **test**: Adding tests
- **chore**: Maintenance tasks

## Code Review

All changes require:
- Passing tests
- Updated documentation
- Code review approval
- CI/CD checks passing
