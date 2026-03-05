package deferexample

// DeferOrder trả về thứ tự thực thi khi dùng defer.
// Append "1", "2", "3" vào slice bằng defer, return slice.
// Kết quả mong đợi: ["3", "2", "1"] (LIFO)
func DeferOrder() []string {
	// TODO: Implement this
	// Hint: dùng named return value và defer để append vào result
	return nil
}

// SafeDivide thực hiện a/b, nếu b == 0 thì recover panic và trả về (0, error).
// Dùng defer + recover pattern.
func SafeDivide(a, b int) (result int, err error) {
	// TODO: Implement this
	return 0, nil
}

// StackTrace mô phỏng push/pop bằng defer.
// Nhận danh sách tên function, "enter" mỗi function theo thứ tự,
// rồi "exit" theo thứ tự ngược (dùng defer).
// Trả về slice string ghi lại thứ tự: ["enter A", "enter B", "exit B", "exit A"]
func StackTrace(names []string) []string {
	// TODO: Implement this
	return nil
}

// CountDown trả về slice đếm ngược từ n đến 1.
// Dùng defer trong loop để tạo thứ tự ngược.
func CountDown(n int) []int {
	// TODO: Implement this
	return nil
}
