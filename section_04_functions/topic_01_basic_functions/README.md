# Topic 01 – Basic Functions

## Cú pháp

```go
func add(a, b int) int {
    return a + b
}
```

## Multiple Return Values

```go
func divide(a, b float64) (float64, error) {
    if b == 0 { return 0, errors.New("division by zero") }
    return a / b, nil
}
```

## Named Returns

```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return // naked return
}
```

## Functions as First-Class Values

```go
op := func(a, b int) int { return a + b }
result := op(3, 5)
```

## Bài tập

Implement các function trong `solution.go`, chạy `go test -v` để kiểm tra.
