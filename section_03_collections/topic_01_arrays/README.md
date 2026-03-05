# Topic 01 – Arrays

## Khai báo

```go
var a [5]int                     // zero-value array
b := [3]string{"Go", "Java", "Python"}
c := [...]int{1, 2, 3}          // compiler đếm length
```

## Đặc điểm

- **Fixed size**: Length là một phần của type. `[3]int` ≠ `[5]int`.
- **Value type**: Gán array = copy toàn bộ.
- **Comparable**: Có thể dùng `==` để so sánh.

## Multi-dimensional

```go
matrix := [2][3]int{
    {1, 2, 3},
    {4, 5, 6},
}
```

## Bài tập

Implement các function trong `solution.go`, chạy `go test -v` để kiểm tra.
