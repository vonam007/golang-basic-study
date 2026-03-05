# Topic 03 – Maps

## Khai báo

```go
m := map[string]int{"a": 1, "b": 2}
m := make(map[string]int)
```

## CRUD Operations

```go
m["key"] = value           // Create / Update
val := m["key"]            // Read (zero value nếu không tồn tại)
val, ok := m["key"]        // Comma-ok idiom
delete(m, "key")           // Delete
```

## Iteration

```go
for key, value := range m { }
// ⚠️ Thứ tự iteration KHÔNG đảm bảo
```

## Zero Value

Map zero value là `nil`. Đọc từ nil map = zero value, nhưng ghi vào nil map = **panic**.

## Bài tập

Implement các function trong `solution.go`, chạy `go test -v` để kiểm tra.
