package channels

// Sum tính tổng các phần tử trong slice và gửi kết quả vào channel.
//
// Params:
//   - nums: slice of integers to sum
//   - ch: send-only channel to deliver the result
//
// Behavior:
//   - Tính tổng tất cả phần tử trong nums
//   - Gửi kết quả vào ch
//   - Nếu nums rỗng, gửi 0
func Sum(nums []int, ch chan<- int) {
	// TODO: Implement this
}

// FanIn gộp hai channel thành một channel duy nhất.
//
// Params:
//   - ch1: receive-only channel of strings
//   - ch2: receive-only channel of strings
//
// Returns:
//   - merged channel chứa tất cả giá trị từ cả ch1 và ch2
//
// Behavior:
//   - Đọc giá trị từ cả ch1 và ch2 đồng thời
//   - Forward tất cả giá trị vào channel trả về
//   - Đóng channel trả về khi cả ch1 và ch2 đều đã đóng
func FanIn(ch1, ch2 <-chan string) <-chan string {
	// TODO: Implement this
	return nil
}

// Pipeline tạo một chuỗi xử lý dữ liệu qua channels.
//
// Stage 1 (Generator): Đẩy từng phần tử của nums vào channel.
// Stage 2 (Square):    Nhận từ stage 1, bình phương giá trị, gửi tiếp.
// Stage 3 (Filter):    Nhận từ stage 2, chỉ giữ lại giá trị > threshold.
//
// Params:
//   - nums: input slice of integers
//   - threshold: giá trị ngưỡng, chỉ giữ lại kết quả > threshold
//
// Returns:
//   - channel chứa các giá trị đã qua pipeline (squared & filtered)
//
// Behavior:
//   - Mỗi stage chạy trong goroutine riêng
//   - Channel trả về sẽ tự đóng khi pipeline hoàn tất
func Pipeline(nums []int, threshold int) <-chan int {
	// TODO: Implement this
	return nil
}
