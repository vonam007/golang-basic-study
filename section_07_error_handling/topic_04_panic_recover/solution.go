package panicrecover

// SafeExecute chạy fn và recover nếu fn panic.
// Trả về error chứa panic message nếu panic xảy ra.
// Trả về nil nếu fn chạy thành công.
func SafeExecute(fn func()) (err error) {
	// TODO: Implement this using defer + recover
	return nil
}

// MustPositive panic nếu n < 0 với message "negative value: <n>".
// Trả về n nếu >= 0.
func MustPositive(n int) int {
	// TODO: Implement this
	return 0
}

// SafePositive gọi MustPositive bên trong SafeExecute.
// Trả về (value, nil) nếu ok, hoặc (0, error) nếu panic.
func SafePositive(n int) (result int, err error) {
	// TODO: Implement this
	return 0, nil
}

// RecoverMiddleware wrap một handler function.
// Nếu handler panic, recover và trả về error message thay vì crash.
func RecoverMiddleware(handler func(string) string) func(string) (string, error) {
	// TODO: Implement this
	return nil
}
