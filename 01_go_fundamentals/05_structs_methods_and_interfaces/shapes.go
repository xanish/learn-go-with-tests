package _5_structs_methods_and_interfaces

import "math"

// Shape makes anything that has an Area and a Perimeter method a shape
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
	Side1  float64
	Side2  float64
}

func (t Triangle) Perimeter() float64 {
	return t.Side1 + t.Side2 + t.Base
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
