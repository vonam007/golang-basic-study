# Topic 03 – Buffered Channels

Buffered channel có capacity → send không block cho đến khi buffer đầy.

```go
ch := make(chan int, 3)
ch <- 1   // không block
ch <- 2   // không block
ch <- 3   // không block
ch <- 4   // BLOCK — buffer đầy
```

## Bài tập

Implement các function trong `solution.go`.
