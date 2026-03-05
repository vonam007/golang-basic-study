package emptyinterface

import "testing"

func TestPrintAll(t *testing.T) {
	got := PrintAll(1, "hello", true)
	want := []string{"1", "hello", "true"}
	if len(got) != len(want) {
		t.Fatalf("len = %d; want %d", len(got), len(want))
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("[%d] = %q; want %q", i, got[i], want[i])
		}
	}
	if len(PrintAll()) != 0 {
		t.Error("PrintAll() should return empty")
	}
}

func TestIsNil(t *testing.T) {
	if !IsNil(nil) {
		t.Error("IsNil(nil) should be true")
	}
	if IsNil(42) {
		t.Error("IsNil(42) should be false")
	}
	if IsNil("") {
		t.Error(`IsNil("") should be false`)
	}
}

func TestTypeCount(t *testing.T) {
	values := []interface{}{1, "a", 2, true, "b", 3.14}
	got := TypeCount(values)
	if got == nil {
		t.Fatal("returned nil")
	}
	if got["int"] != 2 {
		t.Errorf("int = %d; want 2", got["int"])
	}
	if got["string"] != 2 {
		t.Errorf("string = %d; want 2", got["string"])
	}
	if got["bool"] != 1 {
		t.Errorf("bool = %d; want 1", got["bool"])
	}
	if got["other"] != 1 {
		t.Errorf("other = %d; want 1", got["other"])
	}
}

func TestSafeIndex(t *testing.T) {
	arr := []interface{}{"a", 42, true}
	v, ok := SafeIndex(arr, 1)
	if !ok || v != 42 {
		t.Errorf("SafeIndex(1) = (%v, %v)", v, ok)
	}
	_, ok = SafeIndex(arr, 5)
	if ok {
		t.Error("SafeIndex(5) should fail")
	}
	_, ok = SafeIndex(arr, -1)
	if ok {
		t.Error("SafeIndex(-1) should fail")
	}
}
