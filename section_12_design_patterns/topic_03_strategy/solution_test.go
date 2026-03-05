package strategy

import "testing"

func TestBubbleSort(t *testing.T) {
	got := BubbleSort{}.Sort([]int{3, 1, 4, 1, 5})
	want := []int{1, 1, 3, 4, 5}
	for i, v := range got {
		if v != want[i] {
			t.Errorf("[%d] = %d; want %d", i, v, want[i])
		}
	}
}

func TestInsertionSort(t *testing.T) {
	got := InsertionSort{}.Sort([]int{5, 2, 8, 1})
	want := []int{1, 2, 5, 8}
	for i, v := range got {
		if v != want[i] {
			t.Errorf("[%d] = %d; want %d", i, v, want[i])
		}
	}
}

func TestSortContext(t *testing.T) {
	ctx := NewSortContext(BubbleSort{})
	if ctx == nil {
		t.Fatal("nil context")
	}
	got := ctx.Execute([]int{3, 1, 2})
	if got[0] != 1 {
		t.Error("bubble sort failed")
	}

	ctx.SetStrategy(InsertionSort{})
	got = ctx.Execute([]int{9, 7, 8})
	if got[0] != 7 {
		t.Error("insertion sort failed")
	}
}
