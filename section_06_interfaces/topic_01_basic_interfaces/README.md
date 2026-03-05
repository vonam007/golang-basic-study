# Topic 01 – Basic Interfaces

## Implicit Implementation

Go interfaces are satisfied **implicitly** — không cần keyword `implements`:

```go
type Speaker interface { Speak() string }

type Dog struct{ Name string }
func (d Dog) Speak() string { return d.Name + " barks" }
// Dog implements Speaker tự động
```

## Interface Value

```go
var s Speaker = Dog{Name: "Rex"}
s.Speak() // "Rex barks"
```

## Bài tập

Implement các struct và method trong `solution.go`, chạy `go test -v` để kiểm tra.
