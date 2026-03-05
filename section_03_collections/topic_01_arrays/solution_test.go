package arrays

import "testing"

func TestSumArray(t *testing.T) {
	tests := []struct {
		arr  [5]int
		want int
	}{
		{[5]int{1, 2, 3, 4, 5}, 15},
		{[5]int{0, 0, 0, 0, 0}, 0},
		{[5]int{-1, -2, 3, 4, -5}, -1},
	}
	for _, tt := range tests {
		if got := SumArray(tt.arr); got != tt.want {
			t.Errorf("SumArray(%v) = %d; want %d", tt.arr, got, tt.want)
		}
	}
}

func TestMaxInArray(t *testing.T) {
	tests := []struct {
		arr  [5]int
		want int
	}{
		{[5]int{1, 5, 3, 2, 4}, 5},
		{[5]int{-1, -5, -3, -2, -4}, -1},
		{[5]int{7, 7, 7, 7, 7}, 7},
	}
	for _, tt := range tests {
		if got := MaxInArray(tt.arr); got != tt.want {
			t.Errorf("MaxInArray(%v) = %d; want %d", tt.arr, got, tt.want)
		}
	}
}

func TestReverseArray(t *testing.T) {
	tests := []struct {
		arr  [5]int
		want [5]int
	}{
		{[5]int{1, 2, 3, 4, 5}, [5]int{5, 4, 3, 2, 1}},
		{[5]int{1, 1, 1, 1, 1}, [5]int{1, 1, 1, 1, 1}},
	}
	for _, tt := range tests {
		if got := ReverseArray(tt.arr); got != tt.want {
			t.Errorf("ReverseArray(%v) = %v; want %v", tt.arr, got, tt.want)
		}
	}
}

func TestMatrixSum(t *testing.T) {
	m := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	if got := MatrixSum(m); got != 45 {
		t.Errorf("MatrixSum = %d; want 45", got)
	}
	zero := [3][3]int{}
	if got := MatrixSum(zero); got != 0 {
		t.Errorf("MatrixSum(zero) = %d; want 0", got)
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		arr  [5]int
		want bool
	}{
		{[5]int{1, 2, 3, 2, 1}, true},
		{[5]int{1, 2, 3, 4, 5}, false},
		{[5]int{5, 5, 5, 5, 5}, true},
		{[5]int{1, 2, 3, 2, 0}, false},
	}
	for _, tt := range tests {
		if got := IsPalindrome(tt.arr); got != tt.want {
			t.Errorf("IsPalindrome(%v) = %v; want %v", tt.arr, got, tt.want)
		}
	}
}
