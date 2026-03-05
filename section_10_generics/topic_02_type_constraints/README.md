# Topic 02 – Type Constraints

Custom constraints, `~` operator cho underlying types.

```go
type Stringable interface { ~string }
type Ordered interface { ~int | ~float64 | ~string }
```

## Bài tập

Implement trong `solution.go`.
