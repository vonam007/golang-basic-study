# Topic 04 – Panic & Recover

## Panic

`panic("message")` — dừng execution, chạy deferred functions, crash chương trình.

## Recover

`recover()` trong `defer` giúp bắt panic:

```go
func safeDo() (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("recovered: %v", r)
        }
    }()
    panic("boom")
}
```

## Khi nào dùng?

- **Panic**: Chỉ cho unrecoverable errors (programming bugs).
- **Recover**: Middleware, server handlers, testing.

## Bài tập

Implement các function trong `solution.go`.
