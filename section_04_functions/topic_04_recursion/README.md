# Topic 04 – Recursion

## Cú pháp

```go
func factorial(n int) int {
    if n <= 1 { return 1 }
    return n * factorial(n-1)
}
```

## Lưu ý

- Go **không tối ưu** tail recursion → deep recursion có thể stack overflow.
- Luôn có **base case** rõ ràng.
- Cân nhắc dùng iteration cho performance-critical code.

## Bài tập

Implement các function trong `solution.go`, chạy `go test -v` để kiểm tra.
