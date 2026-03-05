package escapeanalysis

import "testing"

func TestNoEscapeSum(t *testing.T) {
	if got := NoEscapeSum([]int{1, 2, 3}); got != 6 {
		t.Errorf("got %d", got)
	}
	if got := NoEscapeSum([]int{}); got != 0 {
		t.Error("empty")
	}
}

func TestEscapePointer(t *testing.T) {
	p := EscapePointer(3, 4)
	if p == nil {
		t.Fatal("nil")
	}
	if p.X != 3 || p.Y != 4 {
		t.Errorf("got %+v", p)
	}
}

func TestAvoidEscape(t *testing.T) {
	p := AvoidEscape(5, 6)
	if p.X != 5 || p.Y != 6 {
		t.Errorf("got %+v", p)
	}
}

func TestSliceGrow(t *testing.T) {
	s, cap := SliceGrow(10)
	if len(s) != 10 {
		t.Errorf("len = %d", len(s))
	}
	if cap < 10 {
		t.Errorf("cap = %d; want >= 10", cap)
	}
	for i, v := range s {
		if v != i {
			t.Errorf("[%d] = %d", i, v)
		}
	}
}
