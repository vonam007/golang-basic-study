# Topic 04 – Defer

## Cú pháp

`defer` schedule một function call chạy **sau khi** function chứa nó return.

```go
func example() {
    defer fmt.Println("world")
    fmt.Println("hello")
}
// Output: hello → world
```

## LIFO Stack

Nhiều `defer` thực thi theo thứ tự **Last In, First Out (LIFO)**:

```go
defer fmt.Println("1")
defer fmt.Println("2")
defer fmt.Println("3")
// Output: 3 → 2 → 1
```

## Argument Evaluation

Arguments của `defer` được **evaluate ngay lập tức**, không phải lúc thực thi:

```go
x := 0
defer fmt.Println(x) // in ra 0, không phải 1
x = 1
```

## Use Cases

1. **Resource cleanup**: Close file, DB connection
2. **Unlock mutex**: `defer mu.Unlock()`
3. **Recover from panic**: `defer func() { recover() }()`

## Bài tập

Implement các function trong `solution.go`, chạy `go test -v` để kiểm tra.
