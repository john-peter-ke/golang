package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoublyLinkedList(t *testing.T) {

	t.Run("Test InsertAtHead", func(t *testing.T) {
		list := NewDoublylyLinkedList[int32]()
		assert.Equal(t, true, list.IsEmpty())
		list.InsertAtHead(48)
		list.InsertAtHead(18)
		list.InsertAtHead(16)
		assert.Equal(t, false, list.IsEmpty())
		assert.Equal(t, 3, list.Size())
	})

	t.Run("Test InsertAtTail", func(t *testing.T) {
		list := NewDoublylyLinkedList[int32]()
		assert.Equal(t, true, list.IsEmpty())
		list.InsertAtTail(48)
		list.InsertAtTail(18)
		list.InsertAtTail(16)
		assert.Equal(t, false, list.IsEmpty())
		assert.Equal(t, 3, list.Size())
	})

	t.Run("Test Delete Node", func(t *testing.T) {
		list := NewDoublylyLinkedList[int32]()
		assert.Equal(t, true, list.IsEmpty())
		assert.Equal(t, 0, list.Size())
		list.InsertAtTail(48)
		list.InsertAtTail(18)
		list.InsertAtTail(16)
		assert.Equal(t, false, list.IsEmpty())
		assert.Equal(t, 3, list.Size())

		list.Delete(48)
		assert.Equal(t, false, list.IsEmpty())
		assert.Equal(t, 2, list.Size())

		list.Delete(18)
		assert.Equal(t, false, list.IsEmpty())
		assert.Equal(t, 1, list.Size())

		list.Delete(110)
		assert.Equal(t, false, list.IsEmpty())
		assert.Equal(t, 1, list.Size())

		list.Delete(16)
		assert.Equal(t, true, list.IsEmpty())
		assert.Equal(t, 0, list.Size())

		list.Delete(16)
		assert.Equal(t, true, list.IsEmpty())
		assert.Equal(t, 0, list.Size())
	})

	t.Run("Test Insert Both", func(t *testing.T) {
		list := NewDoublylyLinkedList[int32]()
		assert.Equal(t, true, list.IsEmpty())
		list.InsertAtHead(18)
		list.InsertAtHead(48)
		list.InsertAtTail(16)
		assert.Equal(t, false, list.IsEmpty())
		assert.Equal(t, 3, list.Size())
	})
}
