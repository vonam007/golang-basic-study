package basicfunctions

import "errors"

// Abs trả về giá trị tuyệt đối.
func Abs(n int) int {
	// TODO: Implement this
	return 0
}

// Max trả về giá trị lớn nhất trong hai số.
func Max(a, b int) int {
	// TODO: Implement this
	return 0
}

// Divide thực hiện phép chia. Trả về error nếu chia cho 0.
func Divide(a, b float64) (float64, error) {
	// TODO: Implement this
	_ = errors.New("placeholder") // hint: dùng errors.New
	return 0, nil
}

// Apply nhận function và hai int, trả về kết quả sau khi áp dụng function.
func Apply(fn func(int, int) int, a, b int) int {
	// TODO: Implement this
	return 0
}

// MinMax trả về (min, max) của một slice int.
// Nếu slice rỗng, trả về (0, 0, false).
func MinMax(nums []int) (min, max int, ok bool) {
	// TODO: Implement this
	return 0, 0, false
}
