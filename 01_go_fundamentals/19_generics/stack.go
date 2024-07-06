package _9_generics

// Stack represents a generic stack that can hold values of any type.
type Stack[T any] struct {
	values []T
}

// Push adds a new value to the top of the stack.
func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

// IsEmpty checks if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}

// Pop removes and returns the top value from the stack.
// It returns (nil, false) if the stack is empty.
func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	index := len(s.values) - 1
	el := s.values[index]
	s.values = s.values[:index]
	return el, true
}
