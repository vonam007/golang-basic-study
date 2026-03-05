package basicerrors

// Divide thực hiện a / b. Trả về error nếu b == 0.
func Divide(a, b float64) (float64, error) {
	// TODO: Implement this
	return 0, nil
}

// ParseAge chuyển string sang int age.
// Trả về error nếu: không phải số, < 0, hoặc > 150.
func ParseAge(s string) (int, error) {
	// TODO: Implement this
	return 0, nil
}

// SafeGet lấy element tại index từ slice.
// Trả về error nếu index < 0 hoặc >= len.
func SafeGet(items []string, index int) (string, error) {
	// TODO: Implement this
	return "", nil
}

// MultiError thực hiện nhiều operations và thu thập tất cả errors.
// Trả về nil nếu không có error, hoặc combined error message.
func MultiError(operations []func() error) error {
	// TODO: Implement this
	return nil
}
