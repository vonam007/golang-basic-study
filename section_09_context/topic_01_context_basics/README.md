# Topic 01 – Context Basics

`context.Context` dùng để truyền deadlines, cancellation signals, và request-scoped values.

```go
ctx := context.Background()     // root context
ctx := context.TODO()           // placeholder
ctx, cancel := context.WithCancel(ctx)
defer cancel()
```

## Bài tập

Implement các function trong `solution.go`.
