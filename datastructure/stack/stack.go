package stack

type stack[T any] struct {
	values []T
}

// New creates and returns a new stack
func NewStack[T any]() *stack[T] {
	return &stack[T]{values: make([]T, 0)}
}

// Size returns the number of items in stack
func (s *stack[T]) Size() int {
	return len(s.values)
}

// IsEmpty returns true if stack is empty
func (s *stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}

// Push adds value to the top of the stack
func (s *stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

// Pop removes top items from the stack and return it
func (s *stack[T]) Pop() T {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue
	}

	index := s.Size() - 1
	var item = s.values[index]
	s.values = s.values[:index]
	return item
}

// Peek returns the top item in stack
func (s *stack[T]) Peek() T {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue
	}

	return s.values[len(s.values)-1]
}
