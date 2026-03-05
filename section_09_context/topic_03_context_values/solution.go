package contextvalues

import "context"

type contextKey string

const (
	UserIDKey    contextKey = "userID"
	RequestIDKey contextKey = "requestID"
)

// WithUserID thêm userID vào context.
func WithUserID(ctx context.Context, userID string) context.Context {
	// TODO: Implement this
	return ctx
}

// GetUserID lấy userID từ context. Trả về ("", false) nếu không có.
func GetUserID(ctx context.Context) (string, bool) {
	// TODO: Implement this
	return "", false
}

// WithRequestID thêm requestID vào context.
func WithRequestID(ctx context.Context, requestID string) context.Context {
	// TODO: Implement this
	return ctx
}

// GetRequestID lấy requestID từ context.
func GetRequestID(ctx context.Context) (string, bool) {
	// TODO: Implement this
	return "", false
}

// LogLine tạo log string theo format "[requestID] userID: message".
// Nếu thiếu field nào dùng "unknown".
func LogLine(ctx context.Context, message string) string {
	// TODO: Implement this
	return ""
}
