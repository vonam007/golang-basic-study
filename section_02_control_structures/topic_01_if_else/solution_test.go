package ifelse

import "testing"

func TestClassify(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{5, "positive"},
		{-3, "negative"},
		{0, "zero"},
		{100, "positive"},
		{-1, "negative"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := Classify(tt.n)
			if got != tt.want {
				t.Errorf("Classify(%d) = %q; want %q", tt.n, got, tt.want)
			}
		})
	}
}

func TestGrade(t *testing.T) {
	tests := []struct {
		score int
		want  string
	}{
		{95, "A"}, {90, "A"},
		{85, "B"}, {80, "B"},
		{75, "C"}, {70, "C"},
		{65, "D"}, {60, "D"},
		{55, "F"}, {0, "F"},
		{-1, "Invalid"}, {101, "Invalid"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := Grade(tt.score)
			if got != tt.want {
				t.Errorf("Grade(%d) = %q; want %q", tt.score, got, tt.want)
			}
		})
	}
}

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{15, "FizzBuzz"},
		{30, "FizzBuzz"},
		{3, "Fizz"},
		{9, "Fizz"},
		{5, "Buzz"},
		{10, "Buzz"},
		{1, "1"},
		{7, "7"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := FizzBuzz(tt.n)
			if got != tt.want {
				t.Errorf("FizzBuzz(%d) = %q; want %q", tt.n, got, tt.want)
			}
		})
	}
}

func TestClampValue(t *testing.T) {
	tests := []struct {
		name     string
		value    int
		min, max int
		want     int
	}{
		{"in range", 5, 1, 10, 5},
		{"below min", -5, 0, 10, 0},
		{"above max", 15, 0, 10, 10},
		{"at min", 0, 0, 10, 0},
		{"at max", 10, 0, 10, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ClampValue(tt.value, tt.min, tt.max)
			if got != tt.want {
				t.Errorf("ClampValue(%d, %d, %d) = %d; want %d",
					tt.value, tt.min, tt.max, got, tt.want)
			}
		})
	}
}
