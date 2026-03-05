# Topic 01 – Variables

## Khai báo biến

```go
// Khai báo đầy đủ
var name string = "Go"

// Type inference
var age = 25

// Short declaration (chỉ dùng trong function)
count := 10

// Multiple declarations
var (
    x int
    y string
    z bool
)
```

## Zero Values

Mỗi kiểu dữ liệu có zero value mặc định:

| Type                      | Zero Value |
| ------------------------- | ---------- |
| `int`, `float64`          | `0`        |
| `string`                  | `""`       |
| `bool`                    | `false`    |
| `pointer`, `slice`, `map` | `nil`      |

## Scope

- **Package-level**: Khai báo ngoài function, accessible trong cùng package.
- **Function-level**: Khai báo trong function, chỉ accessible trong function đó.
- **Block-level**: Khai báo trong `if`, `for`, etc. chỉ accessible trong block đó.

## Bài tập

Implement các function trong `solution.go`, chạy `go test -v` để kiểm tra.
