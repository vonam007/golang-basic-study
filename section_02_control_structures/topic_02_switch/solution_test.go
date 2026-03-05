package switchcase

import (
	"math"
	"testing"
)

func TestDayType(t *testing.T) {
	tests := []struct{ day, want string }{
		{"Monday", "Weekday"}, {"Tuesday", "Weekday"}, {"Wednesday", "Weekday"},
		{"Thursday", "Weekday"}, {"Friday", "Weekday"},
		{"Saturday", "Weekend"}, {"Sunday", "Weekend"},
		{"Holiday", "Invalid"}, {"", "Invalid"},
	}
	for _, tt := range tests {
		t.Run(tt.day, func(t *testing.T) {
			if got := DayType(tt.day); got != tt.want {
				t.Errorf("DayType(%q) = %q; want %q", tt.day, got, tt.want)
			}
		})
	}
}

func TestSeasonFromMonth(t *testing.T) {
	tests := []struct {
		month int
		want  string
	}{
		{12, "Winter"}, {1, "Winter"}, {2, "Winter"},
		{3, "Spring"}, {4, "Spring"}, {5, "Spring"},
		{6, "Summer"}, {7, "Summer"}, {8, "Summer"},
		{9, "Autumn"}, {10, "Autumn"}, {11, "Autumn"},
		{0, "Invalid"}, {13, "Invalid"}, {-1, "Invalid"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := SeasonFromMonth(tt.month); got != tt.want {
				t.Errorf("SeasonFromMonth(%d) = %q; want %q", tt.month, got, tt.want)
			}
		})
	}
}

func TestTypeName(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
		want  string
	}{
		{"int", 42, "int"},
		{"string", "hello", "string"},
		{"bool", true, "bool"},
		{"float64", 3.14, "float64"},
		{"slice", []int{1}, "unknown"},
		{"nil", nil, "unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TypeName(tt.value); got != tt.want {
				t.Errorf("TypeName(%v) = %q; want %q", tt.value, got, tt.want)
			}
		})
	}
}

func TestCalculator(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		operator string
		want     float64
		ok       bool
	}{
		{"add", 10, 3, "add", 13, true},
		{"sub", 10, 3, "sub", 7, true},
		{"mul", 4, 5, "mul", 20, true},
		{"div", 10, 4, "div", 2.5, true},
		{"div by zero", 10, 0, "div", 0, false},
		{"invalid op", 10, 3, "pow", 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := Calculator(tt.a, tt.b, tt.operator)
			if ok != tt.ok {
				t.Errorf("Calculator(%v, %v, %q) ok = %v; want %v", tt.a, tt.b, tt.operator, ok, tt.ok)
				return
			}
			if ok && math.Abs(got-tt.want) > 0.001 {
				t.Errorf("Calculator(%v, %v, %q) = %v; want %v", tt.a, tt.b, tt.operator, got, tt.want)
			}
		})
	}
}
