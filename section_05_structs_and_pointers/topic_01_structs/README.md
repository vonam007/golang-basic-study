# Topic 01 – Structs

## Khai báo

```go
type Person struct {
    Name string
    Age  int
}

p := Person{Name: "Alice", Age: 30}
p2 := Person{"Bob", 25}  // positional (không khuyến khích)
```

## Exported vs Unexported Fields

- **Uppercase** (Name) → exported, accessible từ package khác.
- **Lowercase** (name) → unexported, chỉ trong package.

## Anonymous Struct

```go
point := struct{ X, Y int }{10, 20}
```

## Bài tập

Implement các function trong `solution.go`, chạy `go test -v` để kiểm tra.
