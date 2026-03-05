package basicfunctions

import (
	"math"
	"testing"
)

func TestAbs(t *testing.T) {
	tests := []struct{ n, want int }{{5, 5}, {-5, 5}, {0, 0}}
	for _, tt := range tests {
		if got := Abs(tt.n); got != tt.want {
			t.Errorf("Abs(%d) = %d; want %d", tt.n, got, tt.want)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct{ a, b, want int }{{3, 5, 5}, {5, 3, 5}, {0, 0, 0}, {-1, -5, -1}}
	for _, tt := range tests {
		if got := Max(tt.a, tt.b); got != tt.want {
			t.Errorf("Max(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestDivide(t *testing.T) {
	got, err := Divide(10, 3)
	if err != nil {
		t.Fatalf("Divide(10, 3) unexpected error: %v", err)
	}
	if math.Abs(got-3.3333) > 0.01 {
		t.Errorf("Divide(10, 3) = %v; want ~3.333", got)
	}

	_, err = Divide(10, 0)
	if err == nil {
		t.Error("Divide(10, 0) expected error, got nil")
	}
}

func TestApply(t *testing.T) {
	add := func(a, b int) int { return a + b }
	mul := func(a, b int) int { return a * b }

	if got := Apply(add, 3, 5); got != 8 {
		t.Errorf("Apply(add, 3, 5) = %d; want 8", got)
	}
	if got := Apply(mul, 3, 5); got != 15 {
		t.Errorf("Apply(mul, 3, 5) = %d; want 15", got)
	}
}

func TestMinMax(t *testing.T) {
	min, max, ok := MinMax([]int{3, 1, 4, 1, 5})
	if !ok || min != 1 || max != 5 {
		t.Errorf("MinMax = (%d, %d, %v); want (1, 5, true)", min, max, ok)
	}

	_, _, ok = MinMax([]int{})
	if ok {
		t.Error("MinMax([]) should return ok=false")
	}

	min, max, ok = MinMax([]int{7})
	if !ok || min != 7 || max != 7 {
		t.Errorf("MinMax([7]) = (%d, %d, %v); want (7, 7, true)", min, max, ok)
	}
}
