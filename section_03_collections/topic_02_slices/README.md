# Topic 02 – Slices

## Slice vs Array

- Slice là **reference type**: trỏ đến underlying array.
- Có `len()` (số phần tử hiện tại) và `cap()` (capacity).

```go
s := []int{1, 2, 3}       // slice literal
s := make([]int, 5)        // len=5, cap=5
s := make([]int, 3, 10)    // len=3, cap=10
```

## Append & Copy

```go
s = append(s, 4, 5)        // append elements
s2 := make([]int, len(s))
copy(s2, s)                 // deep copy
```

## Slicing

```go
a := []int{0, 1, 2, 3, 4}
b := a[1:3]   // [1, 2] — shares underlying array
c := a[:2]    // [0, 1]
d := a[3:]    // [3, 4]
```

## Nil vs Empty Slice

```go
var s []int       // nil slice: s == nil, len = 0
s := []int{}      // empty slice: s != nil, len = 0
```

## Bài tập

Implement các function trong `solution.go`, chạy `go test -v` để kiểm tra.
