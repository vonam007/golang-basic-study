# Topic 05 – Decorator

Function wrapping / middleware pattern trong Go.

```go
func withLogging(fn http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Println(r.URL.Path)
        fn(w, r)
    }
}
```

## Bài tập

Implement trong `solution.go`.
