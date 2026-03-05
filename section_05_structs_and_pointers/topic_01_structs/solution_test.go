package structs

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	tests := []struct{ w, h, want float64 }{
		{5, 3, 15}, {10, 10, 100}, {0, 5, 0},
	}
	for _, tt := range tests {
		r := Rectangle{tt.w, tt.h}
		if got := r.Area(); math.Abs(got-tt.want) > 0.01 {
			t.Errorf("Rectangle{%v,%v}.Area() = %v; want %v", tt.w, tt.h, got, tt.want)
		}
	}
}

func TestPerimeter(t *testing.T) {
	r := Rectangle{5, 3}
	if got := r.Perimeter(); math.Abs(got-16) > 0.01 {
		t.Errorf("Perimeter = %v; want 16", got)
	}
}

func TestIsSquare(t *testing.T) {
	if !(Rectangle{5, 5}).IsSquare() {
		t.Error("5x5 should be square")
	}
	if (Rectangle{5, 3}).IsSquare() {
		t.Error("5x3 should not be square")
	}
}

func TestString(t *testing.T) {
	r := Rectangle{5, 3}
	want := "Rectangle(5.00 x 3.00)"
	if got := r.String(); got != want {
		t.Errorf("String() = %q; want %q", got, want)
	}
}

func TestNewRectangle(t *testing.T) {
	r := NewRectangle(5, 3)
	if r.Width != 5 || r.Height != 3 {
		t.Errorf("NewRectangle(5,3) = %+v; want {5,3}", r)
	}
	zero := NewRectangle(-1, 5)
	if zero.Width != 0 || zero.Height != 0 {
		t.Errorf("NewRectangle(-1,5) should return zero Rectangle, got %+v", zero)
	}
}

func TestScale(t *testing.T) {
	r := Rectangle{5, 3}
	scaled := r.Scale(2)
	if scaled.Width != 10 || scaled.Height != 6 {
		t.Errorf("Scale(2) = %+v; want {10,6}", scaled)
	}
}
