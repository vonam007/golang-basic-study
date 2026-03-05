# Topic 01 – Basic Errors

## error interface

```go
type error interface { Error() string }
```

## Tạo error

```go
err := errors.New("something went wrong")
err := fmt.Errorf("failed to process %d", id)
```

## Kiểm tra error

```go
result, err := doSomething()
if err != nil {
    return err  // propagate
}
```

## Bài tập

Implement các function trong `solution.go`, chạy `go test -v` để kiểm tra.
