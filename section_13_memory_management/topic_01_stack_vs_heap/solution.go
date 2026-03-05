package stackvsheap

// StackExample minh hoạ variable trên stack. Trả về giá trị (copy).
func StackExample() int {
	// TODO: Implement this — khai báo local x = 42, return x
	return 0
}

// HeapExample minh hoạ escape to heap. Trả về pointer.
func HeapExample() *int {
	// TODO: Implement this — khai báo local x = 42, return &x
	return nil
}

// NoEscape minh hoạ variable không escape.
// Tính tổng mà không tạo heap allocation.
func NoEscape(a, b int) int {
	// TODO: Implement this
	return 0
}

// EscapeSlice minh hoạ slice escape to heap khi return.
func EscapeSlice(n int) []int {
	// TODO: Implement this — tạo slice size n, fill 0..n-1, return
	return nil
}
