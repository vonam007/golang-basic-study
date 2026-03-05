package customerrors

import "fmt"

// ValidationError chứa field name và message.
type ValidationError struct {
	Field   string
	Message string
}

// Error implements error interface.
func (e *ValidationError) Error() string {
	// TODO: Implement this — format: "<Field>: <Message>"
	_ = fmt.Sprintf("placeholder")
	return ""
}

// NotFoundError chứa resource type và ID.
type NotFoundError struct {
	Resource string
	ID       string
}

// Error implements error interface.
func (e *NotFoundError) Error() string {
	// TODO: Implement this — format: "<Resource> not found: <ID>"
	return ""
}

// ValidateUser kiểm tra user input.
// Trả về *ValidationError nếu name rỗng hoặc age < 0 hoặc age > 150.
func ValidateUser(name string, age int) error {
	// TODO: Implement this
	return nil
}

// FindUser tìm user theo ID trong map.
// Trả về *NotFoundError nếu không tìm thấy.
func FindUser(users map[string]string, id string) (string, error) {
	// TODO: Implement this
	return "", nil
}
