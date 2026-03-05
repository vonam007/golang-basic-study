# Topic 02 – Context Timeout

```go
ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
defer cancel()
```

## Bài tập

Implement các function trong `solution.go`.
