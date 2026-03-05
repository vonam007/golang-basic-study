package recursion

import "testing"

func TestFactorial(t *testing.T) {
	tests := []struct{ n, want int }{
		{0, 1}, {1, 1}, {5, 120}, {10, 3628800}, {-1, -1},
	}
	for _, tt := range tests {
		if got := Factorial(tt.n); got != tt.want {
			t.Errorf("Factorial(%d) = %d; want %d", tt.n, got, tt.want)
		}
	}
}

func TestPower(t *testing.T) {
	tests := []struct{ base, exp, want int }{
		{2, 0, 1}, {2, 1, 2}, {2, 10, 1024}, {3, 3, 27}, {5, 0, 1},
	}
	for _, tt := range tests {
		if got := Power(tt.base, tt.exp); got != tt.want {
			t.Errorf("Power(%d, %d) = %d; want %d", tt.base, tt.exp, got, tt.want)
		}
	}
}

func TestSumDigits(t *testing.T) {
	tests := []struct{ n, want int }{
		{123, 6}, {0, 0}, {9, 9}, {-456, 15}, {1000, 1},
	}
	for _, tt := range tests {
		if got := SumDigits(tt.n); got != tt.want {
			t.Errorf("SumDigits(%d) = %d; want %d", tt.n, got, tt.want)
		}
	}
}

func TestGCD(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{12, 8, 4}, {7, 13, 1}, {100, 25, 25}, {0, 5, 5}, {17, 17, 17},
	}
	for _, tt := range tests {
		if got := GCD(tt.a, tt.b); got != tt.want {
			t.Errorf("GCD(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestFlatten(t *testing.T) {
	input := []interface{}{1, []interface{}{2, []interface{}{3, 4}}, 5}
	got := Flatten(input)
	want := []int{1, 2, 3, 4, 5}
	if len(got) != len(want) {
		t.Fatalf("Flatten len = %d; want %d", len(got), len(want))
	}
	for i, v := range got {
		if v != want[i] {
			t.Errorf("Flatten[%d] = %d; want %d", i, v, want[i])
		}
	}

	empty := Flatten([]interface{}{})
	if len(empty) != 0 {
		t.Errorf("Flatten([]) len = %d; want 0", len(empty))
	}
}
