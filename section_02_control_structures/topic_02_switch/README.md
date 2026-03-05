# Topic 02 – Switch

## Cú pháp

```go
switch day {
case "Mon", "Tue", "Wed", "Thu", "Fri":
    fmt.Println("Weekday")
case "Sat", "Sun":
    fmt.Println("Weekend")
default:
    fmt.Println("Invalid")
}
```

## Switch không cần expression

```go
switch {
case score >= 90:
    return "A"
case score >= 80:
    return "B"
}
```

## Type Switch

```go
switch v := i.(type) {
case int:
    fmt.Println("int:", v)
case string:
    fmt.Println("string:", v)
default:
    fmt.Println("unknown")
}
```

## Fallthrough

Go **không tự động fallthrough** (khác C/Java). Dùng keyword `fallthrough` nếu cần.

## Bài tập

Implement các function trong `solution.go`, chạy `go test -v` để kiểm tra.
