# Topic 03 – Context Values

```go
ctx = context.WithValue(ctx, "userID", "123")
val := ctx.Value("userID").(string)
```

⚠️ Chỉ dùng cho request-scoped data, **không** dùng thay function parameters.

## Bài tập

Implement trong `solution.go`.
