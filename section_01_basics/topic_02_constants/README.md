# Topic 02 – Constants

## Khai báo constant

```go
const Pi = 3.14159
const Greeting string = "Hello"

// Block declaration
const (
    StatusOK    = 200
    StatusError = 500
)
```

## `iota` – Auto-incrementing

```go
const (
    Sunday = iota  // 0
    Monday         // 1
    Tuesday        // 2
)
```

`iota` reset về 0 ở mỗi `const` block mới. Có thể dùng expressions:

```go
const (
    _  = iota             // 0 (skip)
    KB = 1 << (10 * iota) // 1024
    MB                    // 1048576
    GB                    // 1073741824
)
```

## Typed vs Untyped Constants

- **Untyped**: `const x = 5` — linh hoạt, tự adapt kiểu khi dùng.
- **Typed**: `const x int = 5` — cố định kiểu, cần cast khi dùng với kiểu khác.

## Bài tập

Implement các function trong `solution.go`, chạy `go test -v` để kiểm tra.
