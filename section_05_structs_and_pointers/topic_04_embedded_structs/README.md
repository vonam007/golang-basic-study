# Topic 04 – Embedded Structs

## Composition over Inheritance

Go không có inheritance, thay vào đó dùng **embedding**:

```go
type Animal struct { Name string }
func (a Animal) Speak() string { return a.Name + " speaks" }

type Dog struct {
    Animal          // embedded → Dog "inherits" Animal's fields/methods
    Breed string
}

d := Dog{Animal: Animal{Name: "Rex"}, Breed: "Labrador"}
d.Name    // "Rex" — promoted field
d.Speak() // "Rex speaks" — promoted method
```

## Method Override

```go
func (d Dog) Speak() string { return d.Name + " barks" }
```

## Bài tập

Implement các struct/method trong `solution.go`, chạy `go test -v` để kiểm tra.
