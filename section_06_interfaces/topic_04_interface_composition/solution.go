package interfacecomposition

// Saver can save data
type Saver interface {
	Save(data string) error
}

// Loader can load data
type Loader interface {
	Load(key string) (string, error)
}

// Storage combines Saver and Loader
type Storage interface {
	Saver
	Loader
}

// MemoryStore là in-memory implementation của Storage.
type MemoryStore struct {
	data map[string]string
}

// NewMemoryStore tạo MemoryStore mới.
func NewMemoryStore() *MemoryStore {
	// TODO: Implement this
	return nil
}

// Save lưu data vào store với key = data.
func (m *MemoryStore) Save(data string) error {
	// TODO: Implement this
	return nil
}

// Load đọc data từ store theo key.
// Trả về error nếu key không tồn tại.
func (m *MemoryStore) Load(key string) (string, error) {
	// TODO: Implement this
	return "", nil
}

// Deleter interface mở rộng thêm
type Deleter interface {
	Delete(key string) error
}

// FullStorage combines all CRUD operations
type FullStorage interface {
	Storage
	Deleter
}

// Delete xóa key khỏi store. Trả về error nếu key không tồn tại.
func (m *MemoryStore) Delete(key string) error {
	// TODO: Implement this
	return nil
}

// CopyData copy tất cả data từ source sang destination.
// Dùng Loader để đọc keys, Saver để ghi.
func CopyData(src Storage, dst Storage, keys []string) error {
	// TODO: Implement this
	return nil
}
