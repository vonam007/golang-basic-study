# Topic 03 – Empty Interface

`interface{}` (hay `any` trong Go 1.18+) chấp nhận **bất kỳ kiểu nào**.

```go
var x interface{} = 42
x = "hello"
x = true
```

## Use Cases

- Printf-style functions
- JSON/data parsing
- Container cho mixed types

## Hạn chế

- Mất type safety → cần type assertion để dùng.
- Không thể gọi method trực tiếp.

## Bài tập

Implement các function trong `solution.go`, chạy `go test -v` để kiểm tra.
