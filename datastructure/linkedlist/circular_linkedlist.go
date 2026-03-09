package linkedlist

import (
	"fmt"

	"github.com/mucunga90/go/constraint"
)

// Circular Linked Lists & Implementations :
// For Circular Linked Lists, you could simply convert either
// a Singly Linked List or Doubly Linked List by doing this :

func ConvertSinglyToCircular[T constraint.Ordered](l *dlinkedList[T]) {
	p := l.head
	for p.next != nil {
		p = p.next
	}
	p.next = l.head // Link the last node back to the head
}

// and it's traversal would look like this :
func ShowCircular[T constraint.Ordered](l *dlinkedList[T]) {
	if l.head == nil {
		fmt.Println("List is empty")
		return
	}
	p := l.head
	for {
		fmt.Printf("-> %v ", p.data)
		p = p.next
		if p == l.head { // Stop when we return to the head
			break
		}
	}
	fmt.Println()
}
