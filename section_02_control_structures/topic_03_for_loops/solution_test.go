package forloops

import "testing"

func TestSumRange(t *testing.T) {
	tests := []struct{ n, want int }{
		{5, 15}, {10, 55}, {1, 1}, {0, 0}, {-1, 0}, {100, 5050},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := SumRange(tt.n); got != tt.want {
				t.Errorf("SumRange(%d) = %d; want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestReverseString(t *testing.T) {
	tests := []struct{ input, want string }{
		{"Hello", "olleH"},
		{"", ""},
		{"a", "a"},
		{"ab", "ba"},
		{"Gö", "öG"},
		{"12345", "54321"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := ReverseString(tt.input); got != tt.want {
				t.Errorf("ReverseString(%q) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestFindIndex(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"found first", []int{1, 2, 3}, 1, 0},
		{"found mid", []int{1, 2, 3}, 2, 1},
		{"found last", []int{1, 2, 3}, 3, 2},
		{"not found", []int{1, 2, 3}, 5, -1},
		{"empty", []int{}, 1, -1},
		{"duplicates", []int{1, 2, 2, 3}, 2, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindIndex(tt.nums, tt.target); got != tt.want {
				t.Errorf("FindIndex(%v, %d) = %d; want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}

func TestCountVowels(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"Hello World", 3},
		{"AEIOU", 5},
		{"bcdfg", 0},
		{"", 0},
		{"GoLang", 2},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := CountVowels(tt.input); got != tt.want {
				t.Errorf("CountVowels(%q) = %d; want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestFlattenMatrix(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   []int
	}{
		{"2x2", [][]int{{1, 2}, {3, 4}}, []int{1, 2, 3, 4}},
		{"empty", [][]int{}, nil},
		{"single row", [][]int{{1, 2, 3}}, []int{1, 2, 3}},
		{"ragged", [][]int{{1}, {2, 3}, {4, 5, 6}}, []int{1, 2, 3, 4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FlattenMatrix(tt.matrix)
			if len(got) != len(tt.want) {
				t.Fatalf("FlattenMatrix got len %d; want %d", len(got), len(tt.want))
			}
			for i, v := range got {
				if v != tt.want[i] {
					t.Errorf("got[%d] = %d; want %d", i, v, tt.want[i])
				}
			}
		})
	}
}
