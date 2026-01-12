# ADR-004: Use Bun and TypeScript for CLI Tool

**Status**: Accepted

**Date**: 2026-01-11

## Context

We need a fast, modern runtime for our CLI tool that provides:
- Fast startup times
- Built-in TypeScript support
- Easy distribution as a single binary
- Good developer experience
- Cross-platform compatibility

## Decision

We will use Bun runtime with TypeScript for the CLI tool, along with Commander.js for argument parsing.

## Consequences

**Positive**:
- Bun is significantly faster than Node.js
- Native TypeScript support without transpilation
- Can compile to single executable binary
- Modern JavaScript features available
- Built-in test runner
- Better developer experience

**Negative**:
- Bun is newer and less mature than Node.js
- Smaller ecosystem compared to Node.js
- Team needs to install Bun

## Alternatives Considered

- Go: Requires compilation, steeper learning curve
- Node.js: Slower startup, requires separate TypeScript setup
- Deno: Good alternative but less ecosystem support
- Rust: Fast but much steeper learning curve
