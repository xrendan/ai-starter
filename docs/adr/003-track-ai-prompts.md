# ADR-003: Track AI Prompts with git-ai

**Status**: Accepted

**Date**: 2026-01-11

## Context

AI-generated code is becoming a significant portion of modern codebases. We need to:
- Track which code was AI-generated
- Preserve the prompts that generated code
- Understand the evolution of AI interactions
- Maintain accountability and reviewability

Prompts are the new source code - they should be versioned and tracked.

## Decision

We will use git-ai to:
- Automatically track AI-generated code
- Store prompts in git notes
- Score commits by AI contribution
- Preserve AI authorship through git operations

## Consequences

**Positive**:
- Full visibility into AI contributions
- Prompts are preserved for future reference
- Can analyze AI impact on codebase
- Supports multiple AI coding agents

**Negative**:
- Additional tool to install and maintain
- Git notes add some complexity
- Team needs to understand git-ai

## Alternatives Considered

- Manual prompt documentation: Too error-prone
- Comments in code: Clutters codebase
- Separate prompt repository: Loses connection to code
