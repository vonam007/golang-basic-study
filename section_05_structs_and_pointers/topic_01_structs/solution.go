package structs

import "fmt"

// Rectangle đại diện cho hình chữ nhật.
type Rectangle struct {
	Width  float64
	Height float64
}

// Area trả về diện tích hình chữ nhật.
func (r Rectangle) Area() float64 {
	// TODO: Implement this
	_ = fmt.Sprintf("placeholder")
	return 0
}

// Perimeter trả về chu vi hình chữ nhật.
func (r Rectangle) Perimeter() float64 {
	// TODO: Implement this
	return 0
}

// IsSquare kiểm tra có phải hình vuông không.
func (r Rectangle) IsSquare() bool {
	// TODO: Implement this
	return false
}

// String trả về biểu diễn string: "Rectangle(Width x Height)"
func (r Rectangle) String() string {
	// TODO: Implement this using fmt.Sprintf
	return ""
}

// NewRectangle tạo Rectangle mới. Nếu width hoặc height <= 0, trả về zero Rectangle.
func NewRectangle(width, height float64) Rectangle {
	// TODO: Implement this
	return Rectangle{}
}

// Scale nhân width và height với factor, trả về Rectangle mới.
func (r Rectangle) Scale(factor float64) Rectangle {
	// TODO: Implement this
	return Rectangle{}
}
