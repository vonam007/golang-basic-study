package testfixtures

import "os"

// TempFileHelper tạo temp file với content, trả về path.
// Caller phải cleanup sau khi dùng.
func TempFileHelper(content string) (string, error) {
	// TODO: Implement this — dùng os.CreateTemp
	_ = os.TempDir()
	return "", nil
}

// ReadFileContent đọc nội dung file trả về string.
func ReadFileContent(path string) (string, error) {
	// TODO: Implement this
	return "", nil
}

// SetupTestDir tạo temp directory cho testing.
// Trả về (path, cleanup function).
func SetupTestDir() (string, func(), error) {
	// TODO: Implement this
	return "", nil, nil
}
