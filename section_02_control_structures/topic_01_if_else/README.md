# Topic 01 – If/Else

## Cú pháp cơ bản

```go
if x > 0 {
    fmt.Println("positive")
} else if x == 0 {
    fmt.Println("zero")
} else {
    fmt.Println("negative")
}
```

## Init Statement

Go cho phép khai báo biến trong điều kiện `if`:

```go
if val, err := doSomething(); err != nil {
    // handle error
} else {
    // use val
}
// val và err không tồn tại ngoài block này
```

## Lưu ý

- Không cần dấu `()` quanh condition.
- `{` bắt buộc phải cùng dòng với `if`.
- Không có ternary operator (`? :`) trong Go.

## Bài tập

Implement các function trong `solution.go`, chạy `go test -v` để kiểm tra.
