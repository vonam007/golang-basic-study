package variables

import "testing"

func TestSwapValues(t *testing.T) {
	tests := []struct {
		name  string
		a, b  int
		wantA int
		wantB int
	}{
		{"positive", 1, 2, 2, 1},
		{"negative", -5, 10, 10, -5},
		{"same", 7, 7, 7, 7},
		{"zero", 0, 42, 42, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotA, gotB := SwapValues(tt.a, tt.b)
			if gotA != tt.wantA || gotB != tt.wantB {
				t.Errorf("SwapValues(%d, %d) = (%d, %d); want (%d, %d)",
					tt.a, tt.b, gotA, gotB, tt.wantA, tt.wantB)
			}
		})
	}
}

func TestZeroValues(t *testing.T) {
	i, f, s, b := ZeroValues()

	if i != 0 {
		t.Errorf("int zero value = %d; want 0", i)
	}
	if f != 0.0 {
		t.Errorf("float64 zero value = %f; want 0.0", f)
	}
	if s != "" {
		t.Errorf("string zero value = %q; want empty", s)
	}
	if b != false {
		t.Errorf("bool zero value = %v; want false", b)
	}
}

func TestMultiAssign(t *testing.T) {
	tests := []struct {
		name       string
		input      []int
		wantFirst  int
		wantLast   int
		wantLength int
	}{
		{"normal", []int{10, 20, 30}, 10, 30, 3},
		{"single", []int{5}, 5, 5, 1},
		{"empty", []int{}, 0, 0, 0},
		{"two elements", []int{1, 99}, 1, 99, 2},
		{"negative", []int{-3, -1, -7}, -3, -7, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, l, n := MultiAssign(tt.input)
			if f != tt.wantFirst || l != tt.wantLast || n != tt.wantLength {
				t.Errorf("MultiAssign(%v) = (%d, %d, %d); want (%d, %d, %d)",
					tt.input, f, l, n, tt.wantFirst, tt.wantLast, tt.wantLength)
			}
		})
	}
}

func TestShadowExample(t *testing.T) {
	got := ShadowExample()
	want := 10
	if got != want {
		t.Errorf("ShadowExample() = %d; want %d (outer x should not be affected by inner shadow)", got, want)
	}
}
