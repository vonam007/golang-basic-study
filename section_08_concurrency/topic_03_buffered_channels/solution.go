package bufferedchannels

// Producer gửi n items vào buffered channel và đóng nó.
// Mỗi item là giá trị index (0 đến n-1).
func Producer(n int) <-chan int {
	// TODO: Implement this — tạo buffered channel với capacity n
	return nil
}

// BatchProcessor nhận items từ channel, xử lý theo batch (batchSize).
// Trả về slice of slices.
func BatchProcessor(ch <-chan int, batchSize int) [][]int {
	// TODO: Implement this
	return nil
}

// WorkerPool tạo numWorkers goroutines để xử lý jobs từ channel.
// Mỗi worker nhận job từ jobs channel, áp dụng fn, gửi kết quả vào results.
func WorkerPool(jobs []int, numWorkers int, fn func(int) int) []int {
	// TODO: Implement this
	return nil
}
