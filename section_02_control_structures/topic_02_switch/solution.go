package switchcase

// DayType trả về "Weekday" hoặc "Weekend" dựa trên tên ngày (tiếng Anh).
// Input: "Monday", "Tuesday", ..., "Sunday" (case-sensitive).
// Nếu input không hợp lệ, trả về "Invalid".
func DayType(day string) string {
	// TODO: Implement this
	return ""
}

// SeasonFromMonth trả về mùa tương ứng với tháng (1-12).
// 12, 1, 2 → "Winter"
// 3, 4, 5 → "Spring"
// 6, 7, 8 → "Summer"
// 9, 10, 11 → "Autumn"
// Nếu month không hợp lệ → "Invalid"
func SeasonFromMonth(month int) string {
	// TODO: Implement this
	return ""
}

// TypeName trả về tên kiểu dữ liệu của value dùng type switch.
// Hỗ trợ: int → "int", string → "string", bool → "bool", float64 → "float64"
// Các kiểu khác → "unknown"
func TypeName(value interface{}) string {
	// TODO: Implement this
	return ""
}

// Calculator thực hiện phép tính dựa trên operator string.
// Hỗ trợ: "add", "sub", "mul", "div"
// Nếu chia cho 0, trả về (0, false).
// Nếu operator không hợp lệ, trả về (0, false).
func Calculator(a, b float64, operator string) (float64, bool) {
	// TODO: Implement this
	return 0, false
}
