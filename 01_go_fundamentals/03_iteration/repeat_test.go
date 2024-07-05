package _3_iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat character default times", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

	t.Run("repeat character 2 times", func(t *testing.T) {
		repeated := Repeat("a", 2)
		expected := "aa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}
