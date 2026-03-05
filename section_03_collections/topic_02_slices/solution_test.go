package slices

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"with dups", []int{1, 2, 2, 3, 1}, []int{1, 2, 3}},
		{"no dups", []int{1, 2, 3}, []int{1, 2, 3}},
		{"all same", []int{5, 5, 5}, []int{5}},
		{"empty", []int{}, nil},
		{"single", []int{1}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveDuplicates(tt.input)
			if len(got) != len(tt.want) {
				t.Fatalf("RemoveDuplicates(%v) len = %d; want %d", tt.input, len(got), len(tt.want))
			}
			for i, v := range got {
				if v != tt.want[i] {
					t.Errorf("got[%d] = %d; want %d", i, v, tt.want[i])
				}
			}
		})
	}
}

func TestChunk(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		size int
		want [][]int
	}{
		{"even", []int{1, 2, 3, 4}, 2, [][]int{{1, 2}, {3, 4}}},
		{"odd", []int{1, 2, 3, 4, 5}, 2, [][]int{{1, 2}, {3, 4}, {5}}},
		{"size > len", []int{1, 2}, 5, [][]int{{1, 2}}},
		{"size 1", []int{1, 2, 3}, 1, [][]int{{1}, {2}, {3}}},
		{"empty", []int{}, 2, nil},
		{"invalid size", []int{1, 2}, 0, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Chunk(tt.nums, tt.size)
			if len(got) != len(tt.want) {
				t.Fatalf("Chunk len = %d; want %d", len(got), len(tt.want))
			}
			for i, chunk := range got {
				if len(chunk) != len(tt.want[i]) {
					t.Fatalf("chunk[%d] len = %d; want %d", i, len(chunk), len(tt.want[i]))
				}
				for j, v := range chunk {
					if v != tt.want[i][j] {
						t.Errorf("chunk[%d][%d] = %d; want %d", i, j, v, tt.want[i][j])
					}
				}
			}
		})
	}
}

func TestFilter(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }
	isPositive := func(n int) bool { return n > 0 }

	tests := []struct {
		name string
		nums []int
		pred func(int) bool
		want []int
	}{
		{"evens", []int{1, 2, 3, 4, 5}, isEven, []int{2, 4}},
		{"positive", []int{-1, 0, 1, 2}, isPositive, []int{1, 2}},
		{"none match", []int{1, 3, 5}, isEven, nil},
		{"empty", []int{}, isEven, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Filter(tt.nums, tt.pred)
			if len(got) != len(tt.want) {
				t.Fatalf("Filter len = %d; want %d", len(got), len(tt.want))
			}
			for i, v := range got {
				if v != tt.want[i] {
					t.Errorf("got[%d] = %d; want %d", i, v, tt.want[i])
				}
			}
		})
	}
}

func TestMap(t *testing.T) {
	double := func(n int) int { return n * 2 }
	square := func(n int) int { return n * n }

	got1 := Map([]int{1, 2, 3}, double)
	want1 := []int{2, 4, 6}
	for i, v := range got1 {
		if v != want1[i] {
			t.Errorf("Map double[%d] = %d; want %d", i, v, want1[i])
		}
	}

	got2 := Map([]int{1, 2, 3}, square)
	want2 := []int{1, 4, 9}
	for i, v := range got2 {
		if v != want2[i] {
			t.Errorf("Map square[%d] = %d; want %d", i, v, want2[i])
		}
	}
}

func TestReduce(t *testing.T) {
	sum := func(a, b int) int { return a + b }
	product := func(a, b int) int { return a * b }

	if got := Reduce([]int{1, 2, 3, 4}, 0, sum); got != 10 {
		t.Errorf("Reduce sum = %d; want 10", got)
	}
	if got := Reduce([]int{1, 2, 3, 4}, 1, product); got != 24 {
		t.Errorf("Reduce product = %d; want 24", got)
	}
	if got := Reduce([]int{}, 5, sum); got != 5 {
		t.Errorf("Reduce empty = %d; want 5", got)
	}
}
