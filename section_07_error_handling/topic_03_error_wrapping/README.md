# Topic 03 – Error Wrapping

## Wrap errors

```go
err := fmt.Errorf("open config: %w", origErr)
```

## Unwrap

```go
errors.Is(err, ErrNotFound)     // kiểm tra error chain
errors.As(err, &target)          // extract specific error type
errors.Unwrap(err)               // lấy wrapped error
```

## Bài tập

Implement các function trong `solution.go`.
