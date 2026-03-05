# Topic 01 – Stack vs Heap

## Stack

- Allocation nhanh, LIFO, tự cleanup khi function return.
- Dùng cho: local variables, function params, return values.

## Heap

- Quản lý bởi GC, chậm hơn stack.
- Dùng cho: data tồn tại ngoài function scope, pointers shared giữa goroutines.

## Escape Analysis

Go compiler tự quyết định stack vs heap. Kiểm tra bằng:

```bash
go build -gcflags="-m" ./...
```

## Bài tập

Implement trong `solution.go`.
