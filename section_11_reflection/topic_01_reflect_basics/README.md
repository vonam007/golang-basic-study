# Topic 01 – Reflect Basics

`reflect.TypeOf` và `reflect.ValueOf` cho phép inspect types tại runtime.

```go
t := reflect.TypeOf(x)  // Type
v := reflect.ValueOf(x) // Value
v.Kind()                 // Kind (int, string, struct, etc.)
```

## Bài tập

Implement trong `solution.go`.
