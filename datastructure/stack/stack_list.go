package stack

import "container/list"

// stackList structure
type stackList struct {
	list *list.List
}

// New creates and returns Stack of List
func NewStackList() *stackList {
	return &stackList{
		list: list.New(),
	}
}

// Size returns number of items in stack
func (s *stackList) Size() int {
	return s.list.Len()
}

// IsEmpty returns true if stack is empty
func (s *stackList) IsEmpty() bool {
	return s.list.Len() == 0
}

// Push add value to the stack
func (s *stackList) Push(value any) {
	s.list.PushFront(value)
}

// Pop removes value from stack
func (s *stackList) Pop() any {
	if s.IsEmpty() {
		return nil
	}
	item := s.list.Front()
	s.list.Remove(item)
	return item.Value
}

// Peek returns the top value in stack
func (s *stackList) Peek() any {
	if s.IsEmpty() {
		return nil
	}

	return s.list.Front().Value
}
