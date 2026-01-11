# Testing Guide

## Test Philosophy

- Tests are part of every commit
- Tests document expected behavior
- Tests prevent regressions
- Tests enable confident refactoring

## Running Tests

Run all tests:
\'\'\'bash
go test ./...
\'\'\'

Run with coverage:
\'\'\'bash
go test -cover ./...
\'\'\'

Run verbose:
\'\'\'bash
go test -v ./...
\'\'\'

## Writing Tests

### Unit Tests

- Test individual functions and methods
- Use table-driven tests for multiple cases
- Mock external dependencies

### Integration Tests

- Test component interactions
- Use real dependencies when feasible
- Test error handling

### Example

\'\'\'go
func TestExample(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        want    string
        wantErr bool
    }{
        {
            name:  "valid input",
            input: "test",
            want:  "result",
        },
        // Add more cases
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Example(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("unexpected error: %v", err)
            }
            if got != tt.want {
                t.Errorf("got %v, want %v", got, tt.want)
            }
        })
    }
}
\'\'\'

## CI/CD

Tests run automatically on:
- Every push
- Every pull request
- Before deployment

All tests must pass before merging.
