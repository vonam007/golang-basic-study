package strings

import "testing"

func TestIsPalindromeString(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"Racecar", true},
		{"hello", false},
		{"", true},
		{"a", true},
		{"Madam", true},
		{"ab", false},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := IsPalindromeString(tt.s); got != tt.want {
				t.Errorf("IsPalindromeString(%q) = %v; want %v", tt.s, got, tt.want)
			}
		})
	}
}

func TestCamelToSnake(t *testing.T) {
	tests := []struct{ input, want string }{
		{"helloWorld", "hello_world"},
		{"Hello", "hello"},
		{"already", "already"},
		{"myHTTPHandler", "my_h_t_t_p_handler"},
		{"", ""},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := CamelToSnake(tt.input); got != tt.want {
				t.Errorf("CamelToSnake(%q) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestTruncateWithEllipsis(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		maxLen int
		want   string
	}{
		{"truncated", "Hello World", 5, "He..."},
		{"not truncated", "Hi", 5, "Hi"},
		{"exact", "Hello", 5, "Hello"},
		{"min len", "Hello", 3, "..."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TruncateWithEllipsis(tt.s, tt.maxLen); got != tt.want {
				t.Errorf("TruncateWithEllipsis(%q, %d) = %q; want %q",
					tt.s, tt.maxLen, got, tt.want)
			}
		})
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"Hello World", 2},
		{"  spaces  everywhere  ", 2},
		{"single", 1},
		{"", 0},
		{"   ", 0},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := CountWords(tt.s); got != tt.want {
				t.Errorf("CountWords(%q) = %d; want %d", tt.s, got, tt.want)
			}
		})
	}
}

func TestRepeatJoin(t *testing.T) {
	tests := []struct {
		s    string
		n    int
		sep  string
		want string
	}{
		{"go", 3, "-", "go-go-go"},
		{"a", 1, ",", "a"},
		{"x", 0, ",", ""},
		{"hi", 2, "", "hihi"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := RepeatJoin(tt.s, tt.n, tt.sep); got != tt.want {
				t.Errorf("RepeatJoin(%q, %d, %q) = %q; want %q",
					tt.s, tt.n, tt.sep, got, tt.want)
			}
		})
	}
}
