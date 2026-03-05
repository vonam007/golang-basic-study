# Topic 02 – Variadic Functions

## Cú pháp

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

sum(1, 2, 3)        // 6
sum()                // 0

nums := []int{1,2,3}
sum(nums...)         // spread slice
```

## Quy tắc

- Variadic parameter phải là **parameter cuối cùng**.
- Bên trong function, variadic param là một **slice**.
- Spread syntax `...` dùng để truyền slice vào variadic function.

## Bài tập

Implement các function trong `solution.go`, chạy `go test -v` để kiểm tra.
