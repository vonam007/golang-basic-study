# Topic 06 – WaitGroup

`sync.WaitGroup` đợi một nhóm goroutines hoàn thành.

```go
var wg sync.WaitGroup
for i := 0; i < 5; i++ {
    wg.Add(1)
    go func() { defer wg.Done(); doWork() }()
}
wg.Wait() // block cho đến khi counter = 0
```

## Bài tập

Implement các function trong `solution.go`.
