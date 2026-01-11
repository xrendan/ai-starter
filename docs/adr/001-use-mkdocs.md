# ADR-001: Use MkDocs for Documentation

**Status**: Accepted

**Date**: 2026-01-11

## Context

We need a documentation system that:
- Lives in the repository with the code
- Supports versioning alongside code changes
- Is easy to write and maintain
- Generates professional-looking documentation
- Integrates well with CI/CD

## Decision

We will use MkDocs with the Material theme for all project documentation.

## Consequences

**Positive**:
- Documentation lives in the repository
- Markdown is easy to write and review
- Material theme provides excellent UX
- Easy to deploy to GitHub Pages
- Can validate documentation in CI/CD

**Negative**:
- Requires Python for local preview
- Team needs to learn MkDocs configuration

## Alternatives Considered

- GitBook: More features but external hosting
- Jekyll: More complex configuration
- Docusaurus: React-based, heavier weight
