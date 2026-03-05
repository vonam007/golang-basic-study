# Topic 03 – Type Conversions

## Chuyển đổi kiểu cơ bản

Go **không** tự động chuyển đổi kiểu (no implicit conversion). Phải explicit:

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

## String ↔ Number

```go
import "strconv"

// String → Int
n, err := strconv.Atoi("42")

// Int → String
s := strconv.Itoa(42)

// String → Float
f, err := strconv.ParseFloat("3.14", 64)
```

## String ↔ Byte/Rune

```go
s := "Hello"
b := []byte(s)    // string → []byte
s2 := string(b)   // []byte → string

r := []rune("Gö") // string → []rune (Unicode-safe)
```

## Type Alias

```go
type Celsius float64
type Fahrenheit float64

// Cần explicit conversion giữa các type aliases
func CToF(c Celsius) Fahrenheit {
    return Fahrenheit(c*9/5 + 32)
}
```

## Bài tập

Implement các function trong `solution.go`, chạy `go test -v` để kiểm tra.
