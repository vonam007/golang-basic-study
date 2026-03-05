package typeconstraints

import "testing"

func TestSort(t *testing.T) {
	ints := []int{3, 1, 4, 1, 5}
	Sort(ints)
	want := []int{1, 1, 3, 4, 5}
	for i, v := range ints {
		if v != want[i] {
			t.Errorf("[%d] = %d; want %d", i, v, want[i])
		}
	}
	strs := []string{"banana", "apple", "cherry"}
	Sort(strs)
	if strs[0] != "apple" {
		t.Errorf("strings sort: %v", strs)
	}
}

func TestMin(t *testing.T) {
	if Min([]int{3, 1, 4}) != 1 {
		t.Error("Min ints")
	}
	if Min([]float64{3.14, 2.71, 1.41}) != 1.41 {
		t.Error("Min floats")
	}
	if Min([]string{"c", "a", "b"}) != "a" {
		t.Error("Min strings")
	}
}

func TestUnique(t *testing.T) {
	got := Unique([]int{1, 2, 2, 3, 1})
	if len(got) != 3 {
		t.Errorf("Unique len = %d; want 3", len(got))
	}
	gotS := Unique([]string{"a", "b", "a"})
	if len(gotS) != 2 {
		t.Errorf("Unique strings len = %d", len(gotS))
	}
}

func TestWithCustomTypes(t *testing.T) {
	ids := []UserID{3, 1, 2}
	Sort(ids)
	if ids[0] != 1 {
		t.Errorf("UserID sort: %v", ids)
	}
	scores := []Score{9.5, 8.0, 9.8}
	if Min(scores) != 8.0 {
		t.Error("Score min")
	}
}
