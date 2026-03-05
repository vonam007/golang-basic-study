package slices

// RemoveDuplicates loại bỏ phần tử trùng lặp, giữ thứ tự xuất hiện đầu tiên.
// Ví dụ: [1,2,2,3,1] → [1,2,3]
func RemoveDuplicates(nums []int) []int {
	// TODO: Implement this
	return nil
}

// Chunk chia slice thành các nhóm có kích thước size.
// Nhóm cuối có thể nhỏ hơn size.
// Ví dụ: Chunk([1,2,3,4,5], 2) → [[1,2],[3,4],[5]]
// Nếu size <= 0, trả về nil.
func Chunk(nums []int, size int) [][]int {
	// TODO: Implement this
	return nil
}

// Filter trả về slice mới chỉ chứa các phần tử thỏa điều kiện predicate.
func Filter(nums []int, predicate func(int) bool) []int {
	// TODO: Implement this
	return nil
}

// Map áp dụng function transform lên từng phần tử và trả về slice mới.
func Map(nums []int, transform func(int) int) []int {
	// TODO: Implement this
	return nil
}

// Reduce gộp slice thành một giá trị duy nhất.
// Bắt đầu với initial, lần lượt gọi reducer(accumulator, element).
func Reduce(nums []int, initial int, reducer func(int, int) int) int {
	// TODO: Implement this
	return 0
}
