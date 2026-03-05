# Topic 02 – Channels

## Khái niệm

Channel là cơ chế giao tiếp giữa các goroutines trong Go. Tuân theo triết lý:

> _"Don't communicate by sharing memory; share memory by communicating."_

## Cú pháp cơ bản

```go
// Tạo channel
ch := make(chan int)       // unbuffered
ch := make(chan int, 5)    // buffered (capacity 5)

// Gửi và nhận
ch <- 42      // send
val := <-ch   // receive

// Đóng channel
close(ch)
```

## Channel Direction

Giới hạn quyền truy cập của channel qua function parameters:

```go
func send(ch chan<- int) { }  // chỉ gửi (send-only)
func recv(ch <-chan int) { }  // chỉ nhận (receive-only)
```

Compiler sẽ báo lỗi nếu vi phạm direction → tăng type safety.

## Các pattern quan trọng

### 1. Fan-In

Gộp nhiều channel thành một:

```go
func FanIn(ch1, ch2 <-chan string) <-chan string {
    merged := make(chan string)
    go func() { for v := range ch1 { merged <- v } }()
    go func() { for v := range ch2 { merged <- v } }()
    return merged
}
```

### 2. Pipeline

Chuỗi các stage xử lý dữ liệu qua channel:

```
[Generator] → ch1 → [Stage 1] → ch2 → [Stage 2] → ch3 → [Consumer]
```

Mỗi stage là một goroutine, nhận input từ channel trước và gửi output ra channel sau.

## Lưu ý quan trọng

| Tình huống                                     | Kết quả                    |
| ---------------------------------------------- | -------------------------- |
| Send vào unbuffered channel, không có receiver | **Deadlock**               |
| Receive từ channel đã `close()`                | Nhận zero value            |
| Send vào channel đã `close()`                  | **Panic**                  |
| Range over channel                             | Tự dừng khi channel closed |

## Bài tập

Implement 3 functions trong `solution.go`, chạy `go test -v` để kiểm tra kết quả.
