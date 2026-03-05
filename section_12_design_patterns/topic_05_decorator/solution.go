package decorator

// StringFunc là function nhận string trả string.
type StringFunc func(string) string

// WithPrefix trả về decorator thêm prefix.
func WithPrefix(prefix string) func(StringFunc) StringFunc {
	// TODO: Implement this
	return nil
}

// WithSuffix trả về decorator thêm suffix.
func WithSuffix(suffix string) func(StringFunc) StringFunc {
	// TODO: Implement this
	return nil
}

// WithUpperCase trả về decorator chuyển result sang uppercase.
func WithUpperCase() func(StringFunc) StringFunc {
	// TODO: Implement this
	return nil
}

// Chain áp dụng nhiều decorators theo thứ tự.
func Chain(fn StringFunc, decorators ...func(StringFunc) StringFunc) StringFunc {
	// TODO: Implement this
	return nil
}
