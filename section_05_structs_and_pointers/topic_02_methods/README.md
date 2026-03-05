# Topic 02 – Methods

Value receiver vs pointer receiver, method sets.

## Value Receiver

```go
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
```

Nhận **copy** của struct. Không thay đổi original.

## Pointer Receiver

```go
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}
```

Nhận **pointer** → có thể mutate struct gốc.

## Khi nào dùng Pointer Receiver?

1. Cần **mutate** struct.
2. Struct **lớn** → tránh copy.
3. **Consistency**: nếu 1 method dùng pointer receiver, nên dùng cho tất cả.

## Bài tập

Implement các method trong `solution.go`, chạy `go test -v` để kiểm tra.
