package genericds

// Stack là generic LIFO stack.
type Stack[T any] struct {
	items []T
}

// Push thêm item vào top of stack.
func (s *Stack[T]) Push(item T) {
	// TODO: Implement this
}

// Pop lấy item từ top. Trả về (value, true) hoặc (zero, false) nếu rỗng.
func (s *Stack[T]) Pop() (T, bool) {
	// TODO: Implement this
	var zero T
	return zero, false
}

// Peek xem top item mà không remove.
func (s *Stack[T]) Peek() (T, bool) {
	// TODO: Implement this
	var zero T
	return zero, false
}

// Len trả về số phần tử.
func (s *Stack[T]) Len() int {
	// TODO: Implement this
	return 0
}

// Result là generic type cho operation có thể thành công hoặc fail.
type Result[T any] struct {
	Value T
	Err   error
}

// NewSuccess tạo Result thành công.
func NewSuccess[T any](value T) Result[T] {
	// TODO: Implement this
	return Result[T]{}
}

// NewFailure tạo Result thất bại.
func NewFailure[T any](err error) Result[T] {
	// TODO: Implement this
	return Result[T]{}
}

// IsOk kiểm tra Result có thành công không.
func (r Result[T]) IsOk() bool {
	// TODO: Implement this
	return false
}

// Unwrap trả về value. Panic nếu là failure.
func (r Result[T]) Unwrap() T {
	// TODO: Implement this
	var zero T
	return zero
}
