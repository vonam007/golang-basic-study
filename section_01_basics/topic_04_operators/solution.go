package operators

// BasicMath thực hiện phép toán cơ bản giữa a và b.
// op: "+", "-", "*", "/", "%"
// Nếu chia cho 0, trả về (0, false).
// Nếu op không hợp lệ, trả về (0, false).
func BasicMath(a, b int, op string) (int, bool) {
	// TODO: Implement this
	return 0, false
}

// IsBetween kiểm tra value có nằm trong khoảng [min, max] không.
func IsBetween(value, min, max int) bool {
	// TODO: Implement this
	return false
}

// IsEvenAndPositive kiểm tra n có đồng thời là số chẵn VÀ dương không.
func IsEvenAndPositive(n int) bool {
	// TODO: Implement this
	return false
}

// CountSetBits đếm số bit 1 trong biểu diễn nhị phân của n (unsigned).
// Ví dụ: CountSetBits(7) → 3 (binary: 111)
func CountSetBits(n uint) int {
	// TODO: Implement this
	return 0
}

// IsPowerOfTwo kiểm tra n có phải lũy thừa của 2 hay không.
// Hint: Dùng bitwise AND. n & (n-1) == 0 khi n là power of 2.
// n phải > 0.
func IsPowerOfTwo(n uint) bool {
	// TODO: Implement this
	return false
}
