# Topic 02 – Custom Errors

## Custom Error Type

```go
type NotFoundError struct { ID string }
func (e *NotFoundError) Error() string {
    return fmt.Sprintf("not found: %s", e.ID)
}
```

## Bài tập

Implement các custom error types trong `solution.go`.
