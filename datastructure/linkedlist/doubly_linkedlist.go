package linkedlist

import (
	"fmt"

	"github.com/mucunga90/go/constraint"
)

// node struct
type dnode[T constraint.Ordered] struct {
	data T
	next *dnode[T]
	prev *dnode[T]
}

// dlinkedList struct
type dlinkedList[T constraint.Ordered] struct {
	head   *dnode[T]
	length int
}

// NewDoublylyLinkedList create instance of linkedList
func NewDoublylyLinkedList[T constraint.Ordered]() *dlinkedList[T] {
	return &dlinkedList[T]{}
}

// Size returns number of nodes in stack
func (s *dlinkedList[T]) Size() int {
	return s.length
}

// IsEmpty returns true if stack has empty nodes
func (s *dlinkedList[T]) IsEmpty() bool {
	return s.length == 0
}

// InsertAtHead add item to linkedList
func (l *dlinkedList[T]) InsertAtHead(v T) {
	newNode := &dnode[T]{data: v, next: l.head}
	l.head = newNode
	l.head.prev = newNode
	l.length++
}

// InsertAtTail add item to linkedList
func (l *dlinkedList[T]) InsertAtTail(v T) {
	newNode := &dnode[T]{data: v}
	if l.head == nil {
		l.head = newNode
	} else {
		last := l.head
		for last.next != nil {
			last = last.next
		}

		last.next = newNode
		newNode.prev = last
	}
	l.length++
}

// Delete removes item from linkedList
func (l *dlinkedList[T]) Delete(value T) {
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

// PrintData displays linkedList items
func (l dlinkedList[T]) Display() {
	count := 0
	current := l.head
	for current != nil {
		fmt.Printf("Node %d: Value: %v\n", count, current.data)
		current = current.next
		count++
	}
}
