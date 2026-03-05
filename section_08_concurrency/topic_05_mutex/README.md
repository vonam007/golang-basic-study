# Topic 05 – Mutex

`sync.Mutex` bảo vệ shared data khỏi race conditions.

```go
var mu sync.Mutex
mu.Lock()
// critical section
mu.Unlock()
```

`sync.RWMutex` cho phép multiple readers hoặc single writer.

## Bài tập

Implement các struct/method trong `solution.go`.
