package _9_generics

// StackOfInts and StackOfStrings are aliases for the Stack type.
// They represent stacks that can hold integers and strings respectively.
type StackOfInts = Stack
type StackOfStrings = Stack

// Stack represents a generic stack that can hold values of any type.
type Stack struct {
	// Slice to hold values of any type using empty interface
	// Note that this will allow any type to go on stack, so we
	// can get same stack having int, string, bool etc. which is not
	// correct.
	// Using []interface{} also gets rid of any benefits the compiler
	// gives us for type safety making it hard to find bugs.
	// We won't even be able to test it correctly unless we create
	// duplicate functions like AssertStringEqual, AssertIntEqual etc.
	values []interface{}
}

// Push adds a new value to the top of the stack.
func (s *Stack) Push(value interface{}) {
	s.values = append(s.values, value)
}

// IsEmpty checks if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(s.values) == 0
}

// Pop removes and returns the top value from the stack.
// It returns (nil, false) if the stack is empty.
func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}

	index := len(s.values) - 1
	element := s.values[index]
	s.values = s.values[:index]
	return element, true
}
