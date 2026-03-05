# Topic 03 – Closures

## Khái niệm

Closure là function **bắt (capture)** biến từ scope bên ngoài:

```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

c := counter()
c() // 1
c() // 2
```

## Function Factory

```go
func multiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

double := multiplier(2)
double(5) // 10
```

## Bài tập

Implement các function trong `solution.go`, chạy `go test -v` để kiểm tra.
