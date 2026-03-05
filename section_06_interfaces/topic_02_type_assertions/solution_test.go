package typeassertions

import "testing"

func TestDescribe(t *testing.T) {
	tests := []struct {
		input interface{}
		want  string
	}{
		{42, "int: 42"},
		{"hello", "string: hello"},
		{true, "bool: true"},
		{3.14, "unknown"},
		{nil, "unknown"},
	}
	for _, tt := range tests {
		if got := Describe(tt.input); got != tt.want {
			t.Errorf("Describe(%v) = %q; want %q", tt.input, got, tt.want)
		}
	}
}

func TestSafeString(t *testing.T) {
	s, ok := SafeString("hello")
	if !ok || s != "hello" {
		t.Errorf("SafeString(\"hello\") = (%q, %v)", s, ok)
	}
	_, ok = SafeString(42)
	if ok {
		t.Error("SafeString(42) should return false")
	}
	_, ok = SafeString(nil)
	if ok {
		t.Error("SafeString(nil) should return false")
	}
}

func TestSumMixed(t *testing.T) {
	values := []interface{}{1, "skip", 2.5, true, 3, 4.5}
	got := SumMixed(values)
	want := 11.0 // 1 + 2.5 + 3 + 4.5
	if got != want {
		t.Errorf("SumMixed = %v; want %v", got, want)
	}
	if SumMixed([]interface{}{}) != 0 {
		t.Error("SumMixed([]) should be 0")
	}
}

type myStringer struct{ name string }

func (m myStringer) String() string { return "custom:" + m.name }

func TestToString(t *testing.T) {
	if got := ToString("hello"); got != "hello" {
		t.Errorf("ToString(string) = %q; want hello", got)
	}
	if got := ToString(myStringer{"test"}); got != "custom:test" {
		t.Errorf("ToString(Stringer) = %q; want custom:test", got)
	}
	if got := ToString(42); got != "42" {
		t.Errorf("ToString(42) = %q; want 42", got)
	}
}
