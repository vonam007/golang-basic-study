# Topic 01 – Singleton

Thread-safe singleton dùng `sync.Once`:

```go
var instance *Config
var once sync.Once

func GetConfig() *Config {
    once.Do(func() { instance = &Config{} })
    return instance
}
```

## Bài tập

Implement trong `solution.go`.
