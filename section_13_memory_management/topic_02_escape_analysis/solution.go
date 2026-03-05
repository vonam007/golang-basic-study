package escapeanalysis

// NoEscapeSum tính tổng mà variable stays on stack.
func NoEscapeSum(nums []int) int {
	// TODO: Implement this — dùng local accumulator
	return 0
}

// EscapePointer tạo struct và return pointer (escapes to heap).
type Point struct{ X, Y int }

func EscapePointer(x, y int) *Point {
	// TODO: Implement this — return &Point{x, y}
	return nil
}

// AvoidEscape tương tự nhưng return by value (no escape).
func AvoidEscape(x, y int) Point {
	// TODO: Implement this — return Point{x, y}
	return Point{}
}

// SliceGrow demo slice grow behavior và allocation.
// Tạo slice, append n elements, return final slice và cap.
func SliceGrow(n int) ([]int, int) {
	// TODO: Implement this
	return nil, 0
}
