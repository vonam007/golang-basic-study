# Topic 01 – Type Parameters

## Generic Functions (Go 1.18+)

```go
func Max[T int | float64](a, b T) T {
    if a > b { return a }
    return b
}
Max(3, 5)       // T inferred as int
Max(3.14, 2.71) // T inferred as float64
```

## Bài tập

Implement các generic functions trong `solution.go`.
