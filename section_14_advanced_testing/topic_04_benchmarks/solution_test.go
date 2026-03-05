package benchmarks

import (
	"fmt"
	"testing"
)

func TestFibRecursive(t *testing.T) {
	tests := []struct{ n, want int }{{0, 0}, {1, 1}, {10, 55}, {20, 6765}}
	for _, tt := range tests {
		if got := FibRecursive(tt.n); got != tt.want {
			t.Errorf("FibRecursive(%d) = %d; want %d", tt.n, got, tt.want)
		}
	}
}

func TestFibIterative(t *testing.T) {
	tests := []struct{ n, want int }{{0, 0}, {1, 1}, {10, 55}, {20, 6765}}
	for _, tt := range tests {
		if got := FibIterative(tt.n); got != tt.want {
			t.Errorf("FibIterative(%d) = %d; want %d", tt.n, got, tt.want)
		}
	}
}

func TestConcatBoth(t *testing.T) {
	strs := []string{"a", "b", "c"}
	if ConcatLoop(strs) != "abc" {
		t.Error("ConcatLoop")
	}
	if ConcatBuilder(strs) != "abc" {
		t.Error("ConcatBuilder")
	}
}

// Benchmarks — run with: go test -bench=. -benchmem

func BenchmarkConcatLoop(b *testing.B) {
	strs := make([]string, 100)
	for i := range strs {
		strs[i] = fmt.Sprintf("item%d", i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConcatLoop(strs)
	}
}

func BenchmarkConcatBuilder(b *testing.B) {
	strs := make([]string, 100)
	for i := range strs {
		strs[i] = fmt.Sprintf("item%d", i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConcatBuilder(strs)
	}
}

func BenchmarkFibRecursive20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibRecursive(20)
	}
}

func BenchmarkFibIterative20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibIterative(20)
	}
}
