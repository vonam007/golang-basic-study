# Topic 02 – Test Fixtures

`TestMain` chạy setup/teardown cho toàn bộ package.

```go
func TestMain(m *testing.M) {
    setup()
    code := m.Run()
    teardown()
    os.Exit(code)
}
```

## Bài tập

Implement trong `solution.go`.
