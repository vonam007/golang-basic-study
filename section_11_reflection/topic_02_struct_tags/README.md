# Topic 02 – Struct Tags

Struct tags dùng để annotate fields (JSON, DB, validation).

```go
type User struct {
    Name  string `json:"name" validate:"required"`
    Email string `json:"email,omitempty"`
}
```

## Bài tập

Implement trong `solution.go`.
