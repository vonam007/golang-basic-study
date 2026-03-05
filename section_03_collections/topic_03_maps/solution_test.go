package maps

import (
	"sort"
	"testing"
)

func TestWordCount(t *testing.T) {
	got := WordCount([]string{"go", "is", "go", "fast"})
	if got == nil {
		t.Fatal("WordCount returned nil")
	}
	if got["go"] != 2 || got["is"] != 1 || got["fast"] != 1 {
		t.Errorf("WordCount = %v; want go:2 is:1 fast:1", got)
	}

	empty := WordCount([]string{})
	if len(empty) != 0 {
		t.Errorf("WordCount([]) len = %d; want 0", len(empty))
	}
}

func TestInvertMap(t *testing.T) {
	got := InvertMap(map[string]int{"a": 1, "b": 2})
	if got == nil {
		t.Fatal("InvertMap returned nil")
	}
	if got[1] != "a" || got[2] != "b" {
		t.Errorf("InvertMap = %v; want {1:a, 2:b}", got)
	}

	empty := InvertMap(map[string]int{})
	if len(empty) != 0 {
		t.Errorf("InvertMap({}) len = %d; want 0", len(empty))
	}
}

func TestMergeMaps(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 3, "c": 4}
	got := MergeMaps(m1, m2)
	if got == nil {
		t.Fatal("MergeMaps returned nil")
	}
	if got["a"] != 1 || got["b"] != 3 || got["c"] != 4 {
		t.Errorf("MergeMaps = %v; want a:1 b:3 c:4", got)
	}

	single := MergeMaps(map[string]int{"x": 10})
	if single["x"] != 10 {
		t.Errorf("MergeMaps single = %v", single)
	}
}

func TestGroupBy(t *testing.T) {
	got := GroupBy([]string{"go", "is", "fun", "hi"})
	if got == nil {
		t.Fatal("GroupBy returned nil")
	}

	group2 := got[2]
	sort.Strings(group2)
	if len(group2) != 3 {
		t.Errorf("group[2] len = %d; want 3", len(group2))
	}
	if len(got[3]) != 1 || got[3][0] != "fun" {
		t.Errorf("group[3] = %v; want [fun]", got[3])
	}
}

func TestKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	got := Keys(m)
	if len(got) != 3 {
		t.Fatalf("Keys len = %d; want 3", len(got))
	}
	sort.Strings(got)
	want := []string{"a", "b", "c"}
	for i, v := range got {
		if v != want[i] {
			t.Errorf("Keys[%d] = %q; want %q", i, v, want[i])
		}
	}

	empty := Keys(map[string]int{})
	if len(empty) != 0 {
		t.Errorf("Keys({}) len = %d; want 0", len(empty))
	}
}
