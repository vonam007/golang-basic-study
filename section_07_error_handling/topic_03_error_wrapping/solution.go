package errorwrapping

import "errors"

// Sentinel errors
var (
	ErrNotFound     = errors.New("not found")
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden    = errors.New("forbidden")
)

// WrapError wraps err với context message.
// Format: "<context>: <original error>"
func WrapError(context string, err error) error {
	// TODO: Implement this using fmt.Errorf with %w
	return nil
}

// IsNotFound kiểm tra error chain có chứa ErrNotFound không.
func IsNotFound(err error) bool {
	// TODO: Implement this using errors.Is
	return false
}

// GetUser mô phỏng tìm user. Trả về wrapped ErrNotFound nếu id không phải "1".
func GetUser(id string) (string, error) {
	// TODO: Implement this
	return "", nil
}

// GetUserProfile gọi GetUser và wrap error thêm context nếu có.
func GetUserProfile(id string) (string, error) {
	// TODO: Implement this
	return "", nil
}
