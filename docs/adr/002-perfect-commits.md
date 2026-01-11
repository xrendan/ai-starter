# ADR-002: Enforce Perfect Commit Structure

**Status**: Accepted

**Date**: 2026-01-11

## Context

Following Simon Willison's "perfect commit" philosophy, we want each commit to be:
1. A single, focused change (implementation)
2. Accompanied by tests
3. Accompanied by documentation
4. Linked to an issue for context

This makes commits more reviewable, understandable, and maintainable.

## Decision

We will enforce commit structure through:
- Commit message templates
- Git hooks that validate commit messages
- GitHub Actions that verify commits include tests and docs
- Required issue references

## Consequences

**Positive**:
- More maintainable git history
- Easier code review
- Better long-term understanding of changes
- Self-documenting codebase

**Negative**:
- Requires discipline and training
- May slow down initial development
- Not every commit needs to be perfect (typos, etc.)

## Implementation

- Commit template in .gitmessage
- commit-msg hook for validation
- CI checks for test and doc changes
