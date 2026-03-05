# Topic 04 – Benchmarks

Go benchmark functions bắt đầu bằng `Benchmark` và nhận `*testing.B`:

```go
func BenchmarkFib(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Fib(20)
    }
}
```

Chạy: `go test -bench=. -benchmem`

## Bài tập

Implement functions trong `solution.go`, benchmark tests sẵn trong test file.
