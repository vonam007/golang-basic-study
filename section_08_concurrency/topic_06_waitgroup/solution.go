package waitgroup

// ConcurrentFetch mô phỏng fetch nhiều URLs đồng thời.
// Trả về map[url]result. Mỗi result = fetcher(url).
func ConcurrentFetch(urls []string, fetcher func(string) string) map[string]string {
	// TODO: Implement this using sync.WaitGroup
	return nil
}

// ParallelProcess xử lý mỗi item bằng goroutine riêng.
// Đợi tất cả hoàn thành, trả về results (thứ tự giữ nguyên).
func ParallelProcess(items []int, processor func(int) int) []int {
	// TODO: Implement this
	return nil
}

// ErrGroup chạy nhiều functions đồng thời.
// Trả về error đầu tiên gặp được (nếu có), nil nếu tất cả ok.
func ErrGroup(tasks []func() error) error {
	// TODO: Implement this
	return nil
}
