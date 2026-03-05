package typeconstraints

// Ordered constraint cho comparable + orderable types.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64 | ~string
}

// Sort sắp xếp slice in-place (bubble sort).
func Sort[T Ordered](slice []T) {
	// TODO: Implement this
}

// Min trả về giá trị nhỏ nhất trong slice. Panic nếu slice rỗng.
func Min[T Ordered](slice []T) T {
	// TODO: Implement this
	var zero T
	return zero
}

// Unique trả về slice không có phần tử trùng.
func Unique[T comparable](slice []T) []T {
	// TODO: Implement this
	return nil
}

// custom type dùng ~ operator
type UserID int
type Score float64

// Nếu constraints dùng ~int, UserID (underlying int) cũng được chấp nhận.
