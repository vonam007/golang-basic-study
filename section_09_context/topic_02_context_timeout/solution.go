package contexttimeout

import (
	"context"
	"time"
)

// FetchWithTimeout mô phỏng HTTP request với timeout.
// Nếu work hoàn thành trước timeout → trả về (result, nil).
// Nếu timeout → trả về ("", error).
func FetchWithTimeout(ctx context.Context, workDuration time.Duration) (string, error) {
	// TODO: Implement this
	return "", nil
}

// RetryWithTimeout thử lại operation tối đa maxRetries lần.
// Nếu context hết timeout giữa chừng → dừng và trả error.
func RetryWithTimeout(ctx context.Context, operation func() error, maxRetries int) error {
	// TODO: Implement this
	return nil
}
