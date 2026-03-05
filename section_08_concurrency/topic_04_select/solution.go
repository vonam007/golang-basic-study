package selecttopic

import "time"

// FirstResponse trả về giá trị từ channel nào respond trước.
// Timeout sau duration. Trả về ("", false) nếu timeout.
func FirstResponse(ch1, ch2 <-chan string, timeout time.Duration) (string, bool) {
	// TODO: Implement this using select
	return "", false
}

// Merge đọc từ hai channels cho đến khi cả hai đều đóng.
// Trả về tất cả giá trị đã nhận.
func Merge(ch1, ch2 <-chan int) []int {
	// TODO: Implement this using select
	return nil
}

// Ticker tạo giá trị theo interval cho đến khi done signal.
// Trả về số lượng ticks đã phát.
func Ticker(interval time.Duration, done <-chan struct{}) int {
	// TODO: Implement this
	return 0
}
