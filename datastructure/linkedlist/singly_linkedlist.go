package linkedlist

import (
	"fmt"

	"github.com/mucunga90/go/constraint"
)

// node struct
type snode[T constraint.Ordered] struct {
	data T
	next *snode[T]
}

// slinkedList struct
type slinkedList[T constraint.Ordered] struct {
	head   *snode[T]
	length int
}

// NewLinkedList create instance of linkedList
func NewSinglyLinkedList[T constraint.Ordered]() *slinkedList[T] {
	return &slinkedList[T]{}
}

// Size returns number of nodes in stack
func (s *slinkedList[T]) Size() int {
	return s.length
}

// IsEmpty returns true if stack has empty nodes
func (s *slinkedList[T]) IsEmpty() bool {
	return s.length == 0
}

// InsertAtHead add item to linkedList
func (l *slinkedList[T]) InsertAtHead(v T) {
	newNode := &snode[T]{data: v, next: l.head}
	l.head = newNode
	l.length++
}

// InsertAtTail add item to linkedList
func (l *slinkedList[T]) InsertAtTail(v T) {
	newNode := &snode[T]{data: v}
	if l.head == nil {
		l.head = newNode
	} else {
		last := l.head
		for last.next != nil {
			last = last.next
		}

		last.next = newNode
	}
	l.length++
}

// Delete removes item from linkedList
func (l *slinkedList[T]) Delete(value T) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	current := l.head
	for current.next != nil && current.next.data != value {
		current = current.next
	}
	if current.next != nil {
		current.next = current.next.next
		l.length--
	}
}

// Delete removes item from linkedList
func (l *slinkedList[T]) Reverse() {
	if l.length == 0 {
		return
	}

	if l.head.next == nil {
		return
	}

	var prev *snode[T] = nil
	current := l.head
	for current != nil {
		next := current.next
		current.next = prev
		//6. then the prev comes forward
		prev = current
		//7. while the next goes backward
		current = next
	}
	l.head = prev
}

// PrintData displays linkedList items
func (l slinkedList[T]) Display() {
	count := 0
	current := l.head
	for current != nil {
		fmt.Printf("Node %d: Value: %v\n", count, current.data)
		current = current.next
		count++
	}
}
