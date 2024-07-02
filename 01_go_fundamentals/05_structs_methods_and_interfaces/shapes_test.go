package _5_structs_methods_and_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	assertPerimeter := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			// When to Use %g:
			// 1. General Purpose Output: When you want a floating-point number to be printed in the most readable form,
			// without manually deciding between scientific and decimal notation.
			// 2. Adaptive Precision: When you need to control the significant digits rather than the fixed number of
			// decimal places.
			// 3. Clarity in Scientific and Financial Data: When representing numbers where both very large and very
			// small values are common, and you want to ensure readability without overwhelming the user with too many
			// digits or scientific notation.
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		assertPerimeter(t, rectangle, 40.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		assertPerimeter(t, circle, 62.83185307179586)
	})
}

func TestArea(t *testing.T) {
	TestArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		TestArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		TestArea(t, circle, 314.1592653589793)
	})
}
