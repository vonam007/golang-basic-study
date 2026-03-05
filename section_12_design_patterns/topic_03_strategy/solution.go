package strategy

// Sorter interface cho sorting strategies.
type Sorter interface {
	Sort(data []int) []int
}

// BubbleSort implementation.
type BubbleSort struct{}

func (b BubbleSort) Sort(data []int) []int {
	// TODO: Implement bubble sort (return new sorted slice)
	return nil
}

// InsertionSort implementation.
type InsertionSort struct{}

func (i InsertionSort) Sort(data []int) []int {
	// TODO: Implement insertion sort (return new sorted slice)
	return nil
}

// SortContext holds a strategy and uses it.
type SortContext struct {
	strategy Sorter
}

// NewSortContext tạo context với strategy.
func NewSortContext(s Sorter) *SortContext {
	// TODO: Implement this
	return nil
}

// SetStrategy thay đổi strategy tại runtime.
func (c *SortContext) SetStrategy(s Sorter) {
	// TODO: Implement this
}

// Execute thực hiện sort.
func (c *SortContext) Execute(data []int) []int {
	// TODO: Implement this
	return nil
}
