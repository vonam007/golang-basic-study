package emptyinterface

// PrintAll nhận variadic interface{} và trả về slice string biểu diễn giá trị.
// Dùng fmt.Sprintf("%v") cho mỗi giá trị.
func PrintAll(values ...interface{}) []string {
	// TODO: Implement this
	return nil
}

// IsNil kiểm tra giá trị interface có nil không.
func IsNil(v interface{}) bool {
	// TODO: Implement this
	return false
}

// TypeCount đếm số lượng từng kiểu dữ liệu trong slice.
// Trả về map: {"int": n, "string": n, "bool": n, "other": n}
func TypeCount(values []interface{}) map[string]int {
	// TODO: Implement this
	return nil
}

// SafeIndex truy cập element tại index từ []interface{}.
// Trả về (value, true) nếu index hợp lệ, (nil, false) ngược lại.
func SafeIndex(arr []interface{}, index int) (interface{}, bool) {
	// TODO: Implement this
	return nil, false
}
