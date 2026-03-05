package typeparameters

// Number constraint cho int và float64.
type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

// Max trả về giá trị lớn nhất trong hai giá trị.
func Max[T Number](a, b T) T {
	// TODO: Implement this
	return a
}

// Contains kiểm tra slice có chứa target không.
func Contains[T comparable](slice []T, target T) bool {
	// TODO: Implement this
	return false
}

// Map áp dụng transform function lên mỗi phần tử.
func Map[T any, U any](slice []T, fn func(T) U) []U {
	// TODO: Implement this
	return nil
}

// Filter trả về slice chứa các phần tử thỏa predicate.
func Filter[T any](slice []T, predicate func(T) bool) []T {
	// TODO: Implement this
	return nil
}

// Reduce gộp slice thành một giá trị.
func Reduce[T any, U any](slice []T, initial U, fn func(U, T) U) U {
	// TODO: Implement this
	return initial
}
