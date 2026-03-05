# Topic 03 – Mocking

Interface-based mocking trong Go — inject mock implementations cho testing.

```go
type Mailer interface { Send(to, body string) error }
type MockMailer struct { Calls []string }
func (m *MockMailer) Send(to, body string) error {
    m.Calls = append(m.Calls, to)
    return nil
}
```

## Bài tập

Implement service + mock trong `solution.go`.
