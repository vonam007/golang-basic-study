package mutex

import "sync"

// SafeCounter là thread-safe counter.
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

// Increment tăng counter lên 1 (thread-safe).
func (c *SafeCounter) Increment() {
	// TODO: Implement this
}

// Value trả về giá trị hiện tại (thread-safe).
func (c *SafeCounter) Value() int {
	// TODO: Implement this
	return 0
}

// SafeMap là thread-safe map dùng RWMutex.
type SafeMap struct {
	mu   sync.RWMutex
	data map[string]int
}

// NewSafeMap tạo SafeMap mới.
func NewSafeMap() *SafeMap {
	// TODO: Implement this
	return nil
}

// Set ghi giá trị (exclusive lock).
func (m *SafeMap) Set(key string, value int) {
	// TODO: Implement this
}

// Get đọc giá trị (shared read lock). Trả về (value, ok).
func (m *SafeMap) Get(key string) (int, bool) {
	// TODO: Implement this
	return 0, false
}

// Len trả về số lượng entries.
func (m *SafeMap) Len() int {
	// TODO: Implement this
	return 0
}
