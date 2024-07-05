package _5_structs_methods_and_interfaces

import "math"

// Shape defines the interface for any shape that has Area and Perimeter methods. Anything
// which has the Area and Perimeter methods will be allowed as a parameter to the method
// which accepts a Shape.
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle represents a rectangle with width and height.
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter calculates the perimeter of the rectangle.
func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

// Area calculates the area of the rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle represents a circle with a radius.
type Circle struct {
	Radius float64
}

// Perimeter calculates the perimeter of the circle.
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Area calculates the area of the circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle represents a triangle with a base, height, and two sides.
type Triangle struct {
	Base   float64
	Height float64
	Side1  float64
	Side2  float64
}

// Perimeter calculates the perimeter of the triangle.
func (t Triangle) Perimeter() float64 {
	return t.Side1 + t.Side2 + t.Base
}

// Area calculates the area of the triangle.
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
