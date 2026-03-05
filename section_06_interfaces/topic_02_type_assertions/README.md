# Topic 02 – Type Assertions

Type assertion dùng để extract concrete type từ interface value.

## Syntax

```go
var i interface{} = "hello"
s := i.(string)        // panic nếu không phải string
s, ok := i.(string)    // safe: ok = false nếu fail
```

## Type Switch

```go
switch v := i.(type) {
case string: fmt.Println("string:", v)
case int:    fmt.Println("int:", v)
default:     fmt.Println("unknown")
}
```

## Bài tập

Implement các function trong `solution.go`, chạy `go test -v` để kiểm tra.
