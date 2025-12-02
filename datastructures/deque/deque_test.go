package deque

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {

	t.Run("Test IsEmpty", func(t *testing.T) {
		dq := NewDoublyEndedQueue[int]()
		assert.Equal(t, true, dq.IsEmpty())

		dq.EnQueueFront(1)
		assert.Equal(t, false, dq.IsEmpty())
	})

	t.Run("Test Size", func(t *testing.T) {
		dq := NewDoublyEndedQueue[int]()
		assert.Equal(t, 0, dq.Size())

		dq.EnQueueRear(1)
		assert.Equal(t, 1, dq.Size())
	})

	t.Run("Test EnQueueFront", func(t *testing.T) {
		dq := NewDoublyEndedQueue[int]()
		dq.EnQueueFront(5)
		assert.Equal(t, 5, dq.Front())
		assert.Equal(t, 5, dq.Rear())
		dq.EnQueueFront(4)
		assert.Equal(t, 4, dq.Front())
		assert.Equal(t, 5, dq.Rear())
		dq.EnQueueFront(3)
		assert.Equal(t, 3, dq.Front())
		assert.Equal(t, 5, dq.Rear())
		dq.EnQueueFront(2)
		assert.Equal(t, 2, dq.Front())
		assert.Equal(t, 5, dq.Rear())
		dq.EnQueueFront(1)
		assert.Equal(t, 1, dq.Front())
		assert.Equal(t, 5, dq.Rear())
		dq.EnQueueRear(6)
		assert.Equal(t, 1, dq.Front())
		assert.Equal(t, 6, dq.Rear())
	})

	t.Run("Test DeQueueFront", func(t *testing.T) {
		dq := NewDoublyEndedQueue[int]()
		dq.EnQueueFront(3)
		dq.EnQueueFront(2)
		dq.EnQueueFront(1)

		assert.Equal(t, 1, dq.DeQueueFront())
		assert.Equal(t, 2, dq.Front())
	})

	t.Run("Test EnQueueRear", func(t *testing.T) {
		dq := NewDoublyEndedQueue[int]()
		dq.EnQueueRear(1)
		assert.Equal(t, 1, dq.Rear())
		assert.Equal(t, 1, dq.Front())
		dq.EnQueueRear(2)
		assert.Equal(t, 2, dq.Rear())
		assert.Equal(t, 1, dq.Front())
		dq.EnQueueRear(3)
		assert.Equal(t, 3, dq.Rear())
		assert.Equal(t, 1, dq.Front())
		dq.EnQueueFront(1)
		assert.Equal(t, 3, dq.Rear())
		assert.Equal(t, 1, dq.Front())
	})

	t.Run("Test DeQueueRear", func(t *testing.T) {
		dq := NewDoublyEndedQueue[int]()
		dq.EnQueueRear(1)
		dq.EnQueueRear(2)
		dq.EnQueueRear(3)

		assert.Equal(t, 3, dq.DeQueueRear())
		assert.Equal(t, 2, dq.Rear())

		dq.EnQueueFront(1)
		assert.Equal(t, 2, dq.DeQueueRear())
		assert.Equal(t, 1, dq.DeQueueRear())
		assert.Equal(t, 1, dq.DeQueueRear())
		assert.Equal(t, 0, dq.DeQueueRear())
	})
}
