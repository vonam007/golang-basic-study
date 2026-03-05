package operators

import "testing"

func TestBasicMath(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		op   string
		want int
		ok   bool
	}{
		{"add", 5, 3, "+", 8, true},
		{"subtract", 10, 4, "-", 6, true},
		{"multiply", 3, 7, "*", 21, true},
		{"divide", 10, 3, "/", 3, true},
		{"modulo", 10, 3, "%", 1, true},
		{"divide by zero", 5, 0, "/", 0, false},
		{"modulo by zero", 5, 0, "%", 0, false},
		{"invalid op", 5, 3, "^", 0, false},
		{"negative add", -3, -5, "+", -8, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := BasicMath(tt.a, tt.b, tt.op)
			if ok != tt.ok || got != tt.want {
				t.Errorf("BasicMath(%d, %d, %q) = (%d, %v); want (%d, %v)",
					tt.a, tt.b, tt.op, got, ok, tt.want, tt.ok)
			}
		})
	}
}

func TestIsBetween(t *testing.T) {
	tests := []struct {
		name     string
		value    int
		min, max int
		want     bool
	}{
		{"inside", 5, 1, 10, true},
		{"at min", 1, 1, 10, true},
		{"at max", 10, 1, 10, true},
		{"below", 0, 1, 10, false},
		{"above", 11, 1, 10, false},
		{"negative range", -5, -10, -1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsBetween(tt.value, tt.min, tt.max)
			if got != tt.want {
				t.Errorf("IsBetween(%d, %d, %d) = %v; want %v",
					tt.value, tt.min, tt.max, got, tt.want)
			}
		})
	}
}

func TestIsEvenAndPositive(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{4, true},
		{2, true},
		{3, false},
		{-2, false},
		{0, false},
		{1, false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := IsEvenAndPositive(tt.n)
			if got != tt.want {
				t.Errorf("IsEvenAndPositive(%d) = %v; want %v", tt.n, got, tt.want)
			}
		})
	}
}

func TestCountSetBits(t *testing.T) {
	tests := []struct {
		n    uint
		want int
	}{
		{0, 0},
		{1, 1},
		{7, 3},
		{255, 8},
		{1024, 1},
		{15, 4},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := CountSetBits(tt.n)
			if got != tt.want {
				t.Errorf("CountSetBits(%d) = %d; want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestIsPowerOfTwo(t *testing.T) {
	tests := []struct {
		n    uint
		want bool
	}{
		{1, true},
		{2, true},
		{4, true},
		{8, true},
		{16, true},
		{1024, true},
		{0, false},
		{3, false},
		{6, false},
		{100, false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := IsPowerOfTwo(tt.n)
			if got != tt.want {
				t.Errorf("IsPowerOfTwo(%d) = %v; want %v", tt.n, got, tt.want)
			}
		})
	}
}
