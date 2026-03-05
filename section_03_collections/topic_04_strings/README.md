# Topic 04 – Strings

## String Internals

- Strings trong Go là **immutable** byte sequences.
- Dùng `[]rune` để xử lý Unicode đúng cách.

```go
s := "Hello, 世界"
len(s)           // 13 (bytes, không phải characters)
len([]rune(s))   // 9 (runes/characters)
```

## Package `strings`

```go
strings.Contains(s, "sub")
strings.HasPrefix(s, "He")
strings.HasSuffix(s, "lo")
strings.ToUpper(s)
strings.ToLower(s)
strings.Split(s, ",")
strings.Join(parts, "-")
strings.Replace(s, "old", "new", -1)
strings.TrimSpace(s)
```

## `strings.Builder`

Dùng cho việc nối nhiều string hiệu quả (tránh tạo nhiều string tạm):

```go
var b strings.Builder
b.WriteString("Hello")
b.WriteString(" World")
result := b.String()
```

## Bài tập

Implement các function trong `solution.go`, chạy `go test -v` để kiểm tra.
