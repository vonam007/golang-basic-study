package goroutines

import (
	"sort"
	"testing"
)

func TestConcurrentDoubler(t *testing.T) {
	got := ConcurrentDoubler([]int{1, 2, 3, 4, 5})
	want := []int{2, 4, 6, 8, 10}
	if len(got) != len(want) {
		t.Fatalf("len = %d; want %d", len(got), len(want))
	}
	for i, v := range got {
		if v != want[i] {
			t.Errorf("[%d] = %d; want %d", i, v, want[i])
		}
	}
	if len(ConcurrentDoubler([]int{})) != 0 {
		t.Error("empty should return empty")
	}
}

func TestParallelSum(t *testing.T) {
	nums := make([]int, 100)
	for i := range nums {
		nums[i] = i + 1
	}
	got := ParallelSum(nums, 4)
	if got != 5050 {
		t.Errorf("ParallelSum = %d; want 5050", got)
	}
	if ParallelSum([]int{}, 4) != 0 {
		t.Error("empty sum should be 0")
	}
	if ParallelSum([]int{42}, 1) != 42 {
		t.Error("single element")
	}
}

func TestConcurrentMap(t *testing.T) {
	square := func(n int) int { return n * n }
	got := ConcurrentMap([]int{1, 2, 3, 4}, square)
	want := []int{1, 4, 9, 16}
	if len(got) != len(want) {
		t.Fatalf("len = %d", len(got))
	}
	// Sort both to compare (order might not be guaranteed in some impls)
	sortedGot := make([]int, len(got))
	copy(sortedGot, got)
	sort.Ints(sortedGot)
	sort.Ints(want)
	for i, v := range sortedGot {
		if v != want[i] {
			t.Errorf("[%d] = %d; want %d", i, v, want[i])
		}
	}
}
