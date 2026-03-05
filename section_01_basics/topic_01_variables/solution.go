package variables

// SwapValues nhận hai int và trả về chúng theo thứ tự ngược lại.
func SwapValues(a, b int) (int, int) {
	// TODO: Implement this
	return 0, 0
}

// ZeroValues trả về zero values của các kiểu cơ bản.
// Returns: (int, float64, string, bool)
func ZeroValues() (int, float64, string, bool) {
	// TODO: Implement this
	// Khai báo biến KHÔNG gán giá trị, return chúng
	return 0, 0, "", false
}

// MultiAssign nhận một slice int và trả về first, last, length.
// Nếu slice rỗng, trả về (0, 0, 0).
func MultiAssign(nums []int) (first, last, length int) {
	// TODO: Implement this
	return 0, 0, 0
}

// ShadowExample minh hoạ variable shadowing.
// Trả về giá trị của biến x ở outer scope sau khi inner scope shadow nó.
// Outer x = 10, inner block gán x := 20 (shadowed), return outer x.
func ShadowExample() int {
	// TODO: Implement this
	return 0
}
