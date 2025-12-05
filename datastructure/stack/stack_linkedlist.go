package stack

import "fmt"

// node structure
type node[T any] struct {
	Value T
	Next  *node[T]
}

// New creates and returns stack
func NewStackNode[T any]() *stackNode[T] {
	return &stackNode[T]{top: nil, length: 0}
}

// stack structure
type stackNode[T any] struct {
	top    *node[T]
	length int
}

// Size returns number of nodes in stack
func (s *stackNode[T]) Size() int {
	return s.length
}

// IsEmpty returns true if stack has empty nodes
func (s *stackNode[T]) IsEmpty() bool {
	return s.length == 0
}

// Push adds node to the top of the stack
func (s *stackNode[T]) Push(value T) {
	newNode := &node[T]{
		Value: value,
		Next:  s.top,
	}

	s.top = newNode
	s.length++
}

// Peek return the node value on the top of the stack
func (s *stackNode[T]) Peek() T {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue
	}
	return s.top.Value
}

// Pop removes and returns the top value in stack
func (s *stackNode[T]) Pop() T {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue
	}

	value := s.top.Value
	if s.top.Next == nil {
		s.top = nil
	} else {
		s.top.Value = s.top.Next.Value
		s.top.Next = s.top.Next.Next
	}

	s.length--
	return value
}

// Print displays values in console
func (s *stackNode[T]) Print() {
	for e := s.top; e != nil; e = e.Next {
		fmt.Println(e.Value)
	}
}
