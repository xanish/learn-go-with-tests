package _5_structs_methods_and_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	tests := []struct {
		name         string
		shape        Shape
		hasPerimeter float64
	}{
		{name: "Rectangle", shape: Rectangle{10.0, 10.0}, hasPerimeter: 40.0},
		{name: "Circle", shape: Circle{10.0}, hasPerimeter: 62.83185307179586},
	}

	for _, tt := range tests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.hasPerimeter {
				// When to Use %g:
				// 1. General Purpose Output: When you want a floating-point number to be printed in the most readable
				// form, without manually deciding between scientific and decimal notation.
				// 2. Adaptive Precision: When you need to control the significant digits rather than the fixed number
				// of decimal places.
				// 3. Clarity in Scientific and Financial Data: When representing numbers where both very large and very
				// small values are common, and you want to ensure readability without overwhelming the user with too
				// many digits or scientific notation.
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasPerimeter)
			}
		})
	}
}

func TestArea(t *testing.T) {
	tests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
	}

	for _, tt := range tests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
