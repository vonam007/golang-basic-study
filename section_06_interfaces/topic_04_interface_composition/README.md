# Topic 04 – Interface Composition

Embedding interfaces để tạo interfaces phức tạp hơn:

```go
type Reader interface { Read(p []byte) (n int, err error) }
type Writer interface { Write(p []byte) (n int, err error) }
type ReadWriter interface {
    Reader
    Writer
}
```

Bất kỳ type nào implement cả `Read` và `Write` sẽ tự động satisfy `ReadWriter`.

## Bài tập

Implement các struct và interface trong `solution.go`, chạy `go test -v` để kiểm tra.
