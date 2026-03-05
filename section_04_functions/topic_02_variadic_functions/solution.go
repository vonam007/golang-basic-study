package variadicfunctions

// Sum tính tổng các số truyền vào. Trả 0 nếu không có argument.
func Sum(nums ...int) int {
	// TODO: Implement this
	return 0
}

// ConcatWithSep nối các string lại bằng separator.
// Ví dụ: ConcatWithSep("-", "a", "b", "c") → "a-b-c"
// Không có string → trả ""
func ConcatWithSep(sep string, parts ...string) string {
	// TODO: Implement this
	return ""
}

// MinOf trả về giá trị nhỏ nhất. Phải có ít nhất 1 argument.
// Dùng first + rest pattern: MinOf(first int, rest ...int) int
func MinOf(first int, rest ...int) int {
	// TODO: Implement this
	return 0
}

// WrapWith nhận một base string và nhiều wrapper functions.
// Áp dụng từng wrapper lần lượt.
// Ví dụ: WrapWith("text", addBold, addItalic) → "<i><b>text</b></i>"
func WrapWith(s string, wrappers ...func(string) string) string {
	// TODO: Implement this
	return ""
}
