# Topic 03 – Garbage Collection

Go dùng **tricolor mark-and-sweep** GC. Tuning bằng `GOGC`.

- `GOGC=100` (default): GC trigger khi heap gấp đôi live data.
- `GOGC=off`: Tắt GC.

## Bài tập

Implement trong `solution.go` — tập trung vào memory-efficient patterns.
