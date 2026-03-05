package stackvsheap

import "testing"

func TestStackExample(t *testing.T) {
	if got := StackExample(); got != 42 {
		t.Errorf("got %d", got)
	}
}

func TestHeapExample(t *testing.T) {
	p := HeapExample()
	if p == nil {
		t.Fatal("nil")
	}
	if *p != 42 {
		t.Errorf("got %d", *p)
	}
}

func TestNoEscape(t *testing.T) {
	if got := NoEscape(3, 5); got != 8 {
		t.Errorf("got %d", got)
	}
}

func TestEscapeSlice(t *testing.T) {
	got := EscapeSlice(5)
	if len(got) != 5 {
		t.Fatalf("len = %d", len(got))
	}
	for i, v := range got {
		if v != i {
			t.Errorf("[%d] = %d", i, v)
		}
	}
}
