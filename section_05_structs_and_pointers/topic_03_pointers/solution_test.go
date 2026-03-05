package pointers

import "testing"

func TestIncrement(t *testing.T) {
	x := 5
	Increment(&x)
	if x != 6 {
		t.Errorf("After Increment, x = %d; want 6", x)
	}
	Increment(nil) // should not panic
}

func TestSwapPointers(t *testing.T) {
	a, b := 3, 7
	SwapPointers(&a, &b)
	if a != 7 || b != 3 {
		t.Errorf("After swap: a=%d b=%d; want a=7 b=3", a, b)
	}
}

func TestNewInt(t *testing.T) {
	p := NewInt(42)
	if p == nil {
		t.Fatal("NewInt returned nil")
	}
	if *p != 42 {
		t.Errorf("*NewInt(42) = %d; want 42", *p)
	}
}

func TestLen(t *testing.T) {
	if Len(nil) != 0 {
		t.Error("Len(nil) should be 0")
	}
	head := &Node{1, &Node{2, &Node{3, nil}}}
	if got := Len(head); got != 3 {
		t.Errorf("Len = %d; want 3", got)
	}
}

func TestAppend(t *testing.T) {
	head := Append(nil, 1)
	if head == nil || head.Value != 1 {
		t.Fatal("Append(nil, 1) failed")
	}
	head = Append(head, 2)
	head = Append(head, 3)
	got := ToSlice(head)
	want := []int{1, 2, 3}
	if len(got) != len(want) {
		t.Fatalf("len = %d; want %d", len(got), len(want))
	}
	for i, v := range got {
		if v != want[i] {
			t.Errorf("got[%d] = %d; want %d", i, v, want[i])
		}
	}
}

func TestToSlice(t *testing.T) {
	if got := ToSlice(nil); len(got) != 0 {
		t.Errorf("ToSlice(nil) len = %d; want 0", len(got))
	}
}
