package typeassertions

import "fmt"

// Describe trả về mô tả giá trị dùng type switch.
// int → "int: <value>"
// string → "string: <value>"
// bool → "bool: <value>"
// default → "unknown"
func Describe(i interface{}) string {
	// TODO: Implement this
	_ = fmt.Sprintf("placeholder")
	return ""
}

// SafeString extract string từ interface. Trả về ("", false) nếu không phải string.
func SafeString(i interface{}) (string, bool) {
	// TODO: Implement this
	return "", false
}

// SumMixed tính tổng các giá trị số trong slice interface{}.
// Chỉ cộng int và float64. Bỏ qua các kiểu khác.
func SumMixed(values []interface{}) float64 {
	// TODO: Implement this
	return 0
}

// Stringer interface (tương tự fmt.Stringer)
type Stringer interface {
	String() string
}

// ToString chuyển interface thành string.
// Nếu value implement Stringer → gọi String()
// Nếu là string → trả trực tiếp
// Ngược lại → dùng fmt.Sprintf("%v", value)
func ToString(value interface{}) string {
	// TODO: Implement this
	return ""
}
