package basicinterfaces

import (
	"math"
	"testing"
)

func TestCircleArea(t *testing.T) {
	c := Circle{Radius: 5}
	want := math.Pi * 25
	if got := c.Area(); math.Abs(got-want) > 0.01 {
		t.Errorf("Circle Area = %v; want %v", got, want)
	}
}

func TestCirclePerimeter(t *testing.T) {
	c := Circle{Radius: 5}
	want := 2 * math.Pi * 5
	if got := c.Perimeter(); math.Abs(got-want) > 0.01 {
		t.Errorf("Circle Perimeter = %v; want %v", got, want)
	}
}

func TestSquareArea(t *testing.T) {
	s := Square{Side: 4}
	if got := s.Area(); got != 16 {
		t.Errorf("Square Area = %v; want 16", got)
	}
}

func TestSquarePerimeter(t *testing.T) {
	s := Square{Side: 4}
	if got := s.Perimeter(); got != 16 {
		t.Errorf("Square Perimeter = %v; want 16", got)
	}
}

func TestShapeInterface(t *testing.T) {
	var s Shape = Circle{Radius: 3}
	if s.Area() == 0 {
		t.Error("Circle should satisfy Shape interface")
	}
	s = Square{Side: 3}
	if s.Area() == 0 {
		t.Error("Square should satisfy Shape interface")
	}
}

func TestTotalArea(t *testing.T) {
	shapes := []Shape{Circle{Radius: 1}, Square{Side: 2}}
	got := TotalArea(shapes)
	want := math.Pi + 4
	if math.Abs(got-want) > 0.01 {
		t.Errorf("TotalArea = %v; want %v", got, want)
	}
	if TotalArea([]Shape{}) != 0 {
		t.Error("TotalArea([]) should be 0")
	}
}

func TestLargestShape(t *testing.T) {
	c := Circle{Radius: 5}
	s := Square{Side: 3}
	shapes := []Shape{c, s}
	got := LargestShape(shapes)
	if got == nil {
		t.Fatal("LargestShape returned nil")
	}
	if got.Area() != c.Area() {
		t.Error("Circle(r=5) should be largest")
	}
	if LargestShape([]Shape{}) != nil {
		t.Error("LargestShape([]) should return nil")
	}
}
