package variadicfunctions

import "testing"

func TestSum(t *testing.T) {
	if got := Sum(1, 2, 3); got != 6 {
		t.Errorf("Sum(1,2,3) = %d; want 6", got)
	}
	if got := Sum(); got != 0 {
		t.Errorf("Sum() = %d; want 0", got)
	}
	if got := Sum(10); got != 10 {
		t.Errorf("Sum(10) = %d; want 10", got)
	}
	if got := Sum(-1, 1); got != 0 {
		t.Errorf("Sum(-1,1) = %d; want 0", got)
	}

	nums := []int{5, 10, 15}
	if got := Sum(nums...); got != 30 {
		t.Errorf("Sum(spread) = %d; want 30", got)
	}
}

func TestConcatWithSep(t *testing.T) {
	tests := []struct {
		name  string
		sep   string
		parts []string
		want  string
	}{
		{"basic", "-", []string{"a", "b", "c"}, "a-b-c"},
		{"single", "-", []string{"a"}, "a"},
		{"none", "-", []string{}, ""},
		{"empty sep", "", []string{"a", "b"}, "ab"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConcatWithSep(tt.sep, tt.parts...); got != tt.want {
				t.Errorf("ConcatWithSep(%q, %v) = %q; want %q", tt.sep, tt.parts, got, tt.want)
			}
		})
	}
}

func TestMinOf(t *testing.T) {
	if got := MinOf(5); got != 5 {
		t.Errorf("MinOf(5) = %d; want 5", got)
	}
	if got := MinOf(3, 1, 4, 1, 5); got != 1 {
		t.Errorf("MinOf(3,1,4,1,5) = %d; want 1", got)
	}
	if got := MinOf(-1, -5, 0); got != -5 {
		t.Errorf("MinOf(-1,-5,0) = %d; want -5", got)
	}
}

func TestWrapWith(t *testing.T) {
	bold := func(s string) string { return "<b>" + s + "</b>" }
	italic := func(s string) string { return "<i>" + s + "</i>" }

	if got := WrapWith("text", bold); got != "<b>text</b>" {
		t.Errorf("WrapWith bold = %q; want <b>text</b>", got)
	}
	if got := WrapWith("text", bold, italic); got != "<i><b>text</b></i>" {
		t.Errorf("WrapWith bold+italic = %q; want <i><b>text</b></i>", got)
	}
	if got := WrapWith("text"); got != "text" {
		t.Errorf("WrapWith no wrappers = %q; want text", got)
	}
}
