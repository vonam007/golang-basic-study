package tabledriventests

import "testing"

func TestValidate(t *testing.T) {
	tests := []struct {
		name  string
		email string
		want  bool
	}{
		{"valid", "user@example.com", true},
		{"no at", "userexample.com", false},
		{"no dot", "user@example", false},
		{"empty", "", false},
		{"just at", "@", false},
		{"valid gmail", "test@gmail.com", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Validate(tt.email); got != tt.want {
				t.Errorf("Validate(%q) = %v; want %v", tt.email, got, tt.want)
			}
		})
	}
}

func TestCategorize(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{1, "positive"}, {-1, "negative"}, {0, "zero"}, {100, "positive"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := Categorize(tt.n); got != tt.want {
				t.Errorf("Categorize(%d) = %q; want %q", tt.n, got, tt.want)
			}
		})
	}
}

func TestTruncate(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		maxLen int
		want   string
	}{
		{"no truncate", "hello", 10, "hello"},
		{"truncate", "hello world foo", 11, "hello world"},
		{"short", "hi", 2, "hi"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Truncate(tt.s, tt.maxLen)
			if len(got) > tt.maxLen+3 { // +3 for potential "..."
				t.Errorf("Truncate(%q, %d) too long: %q", tt.s, tt.maxLen, got)
			}
		})
	}
}
