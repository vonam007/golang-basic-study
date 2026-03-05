package contextbasics

import "context"

// DoWorkWithCancel thực hiện work cho đến khi context bị cancel.
// Trả về số iterations đã hoàn thành.
func DoWorkWithCancel(ctx context.Context) int {
	// TODO: Implement this — loop, check ctx.Done() mỗi iteration
	return 0
}

// PropagateContext mô phỏng call chain: A → B → C.
// Nếu context bị cancel ở bất kỳ level nào, return error.
func PropagateContext(ctx context.Context) (string, error) {
	// TODO: Implement this
	return "", nil
}

// CancelAfterN tạo context cancel sau n operations.
// Trả về results đã thu thập trước khi cancel.
func CancelAfterN(items []string, n int) []string {
	// TODO: Implement this using context.WithCancel
	return nil
}
