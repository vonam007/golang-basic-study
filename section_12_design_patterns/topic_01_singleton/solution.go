package singleton

import "sync"

// Database là singleton database connection.
type Database struct {
	Connected bool
	Host      string
}

var (
	instance *Database
	once     sync.Once
)

// GetDatabase trả về singleton Database instance.
// Luôn trả về cùng một instance bất kể gọi bao nhiêu lần.
func GetDatabase() *Database {
	// TODO: Implement this using sync.Once
	return nil
}

// ResetForTesting reset singleton (chỉ dùng trong tests).
func ResetForTesting() {
	instance = nil
	once = sync.Once{}
}
