# Topic 04 – Select

`select` cho phép goroutine đợi trên nhiều channel operations đồng thời.

```go
select {
case msg := <-ch1: // ...
case msg := <-ch2: // ...
case <-time.After(time.Second): // timeout
default: // non-blocking
}
```

## Bài tập

Implement các function trong `solution.go`.
