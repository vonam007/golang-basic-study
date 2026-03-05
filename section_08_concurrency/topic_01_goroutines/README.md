# Topic 01 – Goroutines

## Goroutine

```go
go doSomething()           // launch goroutine
go func() { ... }()       // anonymous goroutine
```

Goroutine là **lightweight thread** do Go runtime quản lý (~2KB stack).

## Scheduling

Go dùng **M:N scheduling**: M goroutines chạy trên N OS threads. `GOMAXPROCS` set số OS threads.

## Lưu ý

- Goroutine **không có** return value → dùng channel để trả kết quả.
- Main goroutine kết thúc → tất cả goroutine con bị kill.

## Bài tập

Implement các function trong `solution.go`.
