package basicinterfaces

import "math"

// Shape interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle struct
type Circle struct {
	Radius float64
}

// Area of Circle. Công thức: π * r²
func (c Circle) Area() float64 {
	// TODO: Implement this
	_ = math.Pi
	return 0
}

// Perimeter of Circle. Công thức: 2 * π * r
func (c Circle) Perimeter() float64 {
	// TODO: Implement this
	return 0
}

// Square struct
type Square struct {
	Side float64
}

// Area of Square
func (s Square) Area() float64 {
	// TODO: Implement this
	return 0
}

// Perimeter of Square
func (s Square) Perimeter() float64 {
	// TODO: Implement this
	return 0
}

// TotalArea tính tổng diện tích của tất cả shapes.
func TotalArea(shapes []Shape) float64 {
	// TODO: Implement this
	return 0
}

// LargestShape trả về shape có diện tích lớn nhất.
// Nếu slice rỗng, trả về nil.
func LargestShape(shapes []Shape) Shape {
	// TODO: Implement this
	return nil
}
