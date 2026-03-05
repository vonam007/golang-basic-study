# Topic 02 – Escape Analysis

Kiểm tra allocation bằng: `go build -gcflags="-m" ./...`

## Rules

- Variable **escape** nếu: address lấy và return/stored ngoài scope.
- Variable **stays on stack** nếu: chỉ dùng locally.

## Bài tập

Implement trong `solution.go` — mỗi function demo một scenario.
