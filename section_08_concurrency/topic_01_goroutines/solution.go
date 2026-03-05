package goroutines

import "sync"

// ConcurrentDoubler nhận slice int, dùng goroutine để nhân đôi từng phần tử.
// Trả về slice mới chứa kết quả (thứ tự giữ nguyên).
func ConcurrentDoubler(nums []int) []int {
	// TODO: Implement this using goroutines + sync
	_ = sync.WaitGroup{}
	return nil
}

// ParallelSum chia slice thành n parts, tính tổng mỗi part bằng goroutine riêng.
// Trả về tổng tất cả.
func ParallelSum(nums []int, numWorkers int) int {
	// TODO: Implement this
	return 0
}

// ConcurrentMap áp dụng function transform lên mỗi phần tử song song.
// Trả về slice kết quả giữ nguyên thứ tự.
func ConcurrentMap(nums []int, transform func(int) int) []int {
	// TODO: Implement this
	return nil
}
