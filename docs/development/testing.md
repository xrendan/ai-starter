# Testing Guide

## Test Philosophy

- Tests are part of every commit
- Tests document expected behavior
- Tests prevent regressions
- Tests enable confident refactoring

## Running Tests

Run all tests:
```bash
bun test
```

Run with watch mode:
```bash
bun test --watch
```

## Writing Tests

### Unit Tests

- Test individual functions and methods
- Use Bun's test framework
- Mock external dependencies

### Integration Tests

- Test component interactions
- Use real dependencies when feasible
- Test error handling

### Example

```typescript
import { describe, test, expect } from 'bun:test';

describe('example', () => {
  test('should work correctly', () => {
    const result = exampleFunction('input');
    expect(result).toBe('expected');
  });
});
```

## CI/CD

Tests run automatically on:
- Every push
- Every pull request
- Before deployment

All tests must pass before merging.
