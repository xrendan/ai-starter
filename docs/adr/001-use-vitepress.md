# ADR-001: Use VitePress for Documentation

**Status**: Accepted

**Date**: 2026-01-12

## Context

We need a documentation system that:
- Lives in the repository with the code
- Supports versioning alongside code changes
- Is easy to write and maintain
- Generates professional-looking documentation
- Integrates well with CI/CD
- Uses TypeScript/JavaScript (matching our tech stack)

## Decision

We will use VitePress for all project documentation.

## Consequences

**Positive**:
- Documentation lives in the repository
- Markdown is easy to write and review
- Fast, Vue-powered static site generation
- Easy to deploy to GitHub Pages
- TypeScript-based (matches our stack)
- Built-in search functionality
- Excellent developer experience
- No Python dependency required

**Negative**:
- Requires Node.js/Bun for local preview
- Team needs to learn VitePress configuration
- Newer tool with smaller ecosystem than MkDocs

## Alternatives Considered

- MkDocs: Python-based, doesn't match our TypeScript stack
- Docusaurus: React-based, heavier and more complex
- GitBook: More features but external hosting
- Nextra: Good alternative but less mature
