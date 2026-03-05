# Topic 03 – Pointers

## Cú pháp

```go
x := 42
p := &x     // p là *int, trỏ đến x
*p = 100    // dereference: x giờ = 100
```

## Pointer to Struct

```go
p := &Person{Name: "Alice"}
p.Name = "Bob"  // tự động dereference
```

## Nil Pointer

```go
var p *int  // nil
*p          // panic: nil pointer dereference
```

## Không có Pointer Arithmetic

Go **không hỗ trợ** pointer arithmetic (`p++`, `p+1`) — an toàn hơn C/C++.

## Bài tập

Implement các function trong `solution.go`, chạy `go test -v` để kiểm tra.
