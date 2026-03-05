package typeparameters

import "testing"

func TestMax(t *testing.T) {
	if Max(3, 5) != 5 {
		t.Error("Max(3,5)")
	}
	if Max(5, 3) != 5 {
		t.Error("Max(5,3)")
	}
	if Max(3.14, 2.71) != 3.14 {
		t.Error("Max(3.14,2.71)")
	}
}

func TestContains(t *testing.T) {
	if !Contains([]int{1, 2, 3}, 2) {
		t.Error("should find 2")
	}
	if Contains([]int{1, 2, 3}, 5) {
		t.Error("should not find 5")
	}
	if !Contains([]string{"a", "b"}, "a") {
		t.Error("should find a")
	}
	if Contains([]string{}, "a") {
		t.Error("empty slice")
	}
}

func TestMap(t *testing.T) {
	got := Map([]int{1, 2, 3}, func(n int) string { return string(rune('0' + n)) })
	want := []string{"1", "2", "3"}
	for i, v := range got {
		if v != want[i] {
			t.Errorf("[%d] = %q; want %q", i, v, want[i])
		}
	}
}

func TestFilter(t *testing.T) {
	got := Filter([]int{1, 2, 3, 4, 5}, func(n int) bool { return n%2 == 0 })
	if len(got) != 2 || got[0] != 2 || got[1] != 4 {
		t.Errorf("Filter evens = %v", got)
	}
}

func TestReduce(t *testing.T) {
	sum := Reduce([]int{1, 2, 3, 4}, 0, func(acc, n int) int { return acc + n })
	if sum != 10 {
		t.Errorf("Reduce sum = %d; want 10", sum)
	}
	concat := Reduce([]string{"a", "b", "c"}, "", func(acc string, s string) string { return acc + s })
	if concat != "abc" {
		t.Errorf("Reduce concat = %q", concat)
	}
}
