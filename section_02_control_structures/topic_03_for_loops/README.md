# Topic 03 – For Loops

## Cú pháp

Go chỉ có **một** loại loop: `for`.

```go
// Classic for
for i := 0; i < 10; i++ { }

// While-style
for condition { }

// Infinite loop
for { }

// Range over slice
for index, value := range slice { }

// Range over map
for key, value := range myMap { }

// Range over string (by rune)
for i, ch := range "Hello" { }
```

## Break & Continue

```go
for i := 0; i < 100; i++ {
    if i == 50 { break }     // thoát loop
    if i%2 == 0 { continue } // skip iteration
}
```

## Labeled break

```go
outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if i == 1 && j == 1 { break outer }
    }
}
```

## Bài tập

Implement các function trong `solution.go`, chạy `go test -v` để kiểm tra.
