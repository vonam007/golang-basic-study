package pointers

// Increment tăng giá trị mà pointer trỏ đến lên 1.
// Nếu pointer nil, không làm gì.
func Increment(p *int) {
	// TODO: Implement this
}

// SwapPointers hoán đổi giá trị của hai biến thông qua pointers.
func SwapPointers(a, b *int) {
	// TODO: Implement this
}

// NewInt tạo và trả về pointer đến giá trị int mới.
func NewInt(val int) *int {
	// TODO: Implement this
	return nil
}

// LinkedList node
type Node struct {
	Value int
	Next  *Node
}

// Len trả về độ dài linked list.
func Len(head *Node) int {
	// TODO: Implement this
	return 0
}

// Append thêm giá trị vào cuối linked list. Trả về head.
// Nếu head nil, tạo node mới và trả về.
func Append(head *Node, val int) *Node {
	// TODO: Implement this
	return nil
}

// ToSlice chuyển linked list thành slice.
func ToSlice(head *Node) []int {
	// TODO: Implement this
	return nil
}
