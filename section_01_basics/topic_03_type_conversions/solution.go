package typeconversions

// Celsius và Fahrenheit là custom types cho nhiệt độ.
type Celsius float64
type Fahrenheit float64

// CelsiusToFahrenheit chuyển đổi Celsius sang Fahrenheit.
// Công thức: F = C * 9/5 + 32
func CelsiusToFahrenheit(c Celsius) Fahrenheit {
	// TODO: Implement this
	return 0
}

// FahrenheitToCelsius chuyển đổi Fahrenheit sang Celsius.
// Công thức: C = (F - 32) * 5/9
func FahrenheitToCelsius(f Fahrenheit) Celsius {
	// TODO: Implement this
	return 0
}

// StringToInt chuyển string sang int.
// Trả về (giá trị, nil) nếu thành công.
// Trả về (0, error) nếu string không phải số hợp lệ.
func StringToInt(s string) (int, error) {
	// TODO: Implement this using strconv.Atoi
	return 0, nil
}

// IntToString chuyển int sang string.
func IntToString(n int) string {
	// TODO: Implement this using strconv.Itoa
	return ""
}

// ByteCountToHuman chuyển đổi số bytes sang string human-readable.
// Ví dụ: 1024 → "1.00 KB", 1048576 → "1.00 MB"
// Hỗ trợ: B, KB, MB, GB. Format: "%.2f <unit>"
func ByteCountToHuman(bytes uint64) string {
	// TODO: Implement this
	return ""
}
