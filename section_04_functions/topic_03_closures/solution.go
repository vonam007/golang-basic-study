package closures

// Counter trả về một function mà mỗi lần gọi sẽ tăng bộ đếm lên 1 và trả về giá trị hiện tại.
// Bắt đầu từ 0: lần gọi 1 → 1, lần gọi 2 → 2, ...
func Counter() func() int {
	// TODO: Implement this
	return nil
}

// Adder trả về function cộng thêm x vào bất kỳ số nào.
// Ví dụ: add5 := Adder(5); add5(3) → 8
func Adder(x int) func(int) int {
	// TODO: Implement this
	return nil
}

// Fibonacci trả về function mà mỗi lần gọi trả về số Fibonacci tiếp theo.
// Sequence: 0, 1, 1, 2, 3, 5, 8, ...
func Fibonacci() func() int {
	// TODO: Implement this
	return nil
}

// Accumulator trả về function nhận int và trả về tổng tích lũy.
// Ví dụ: acc := Accumulator(); acc(5) → 5; acc(3) → 8; acc(-2) → 6
func Accumulator() func(int) int {
	// TODO: Implement this
	return nil
}
