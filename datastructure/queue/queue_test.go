package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {

	t.Run("Test EnQueue", func(t *testing.T) {
		queue := New[int]()
		queue.EnQueue(1)
		queue.EnQueue(2)
		queue.EnQueue(3)
		queue.EnQueue(4)
		queue.EnQueue(5)

		assert.Equal(t, 1, queue.Front())
		assert.Equal(t, 5, queue.Rear())
	})

	t.Run("Test DeQueue", func(t *testing.T) {
		queue := New[int]()
		queue.EnQueue(6)
		queue.EnQueue(7)
		queue.EnQueue(8)
		queue.EnQueue(9)

		assert.Equal(t, 6, queue.DeQueue())
		assert.Equal(t, 7, queue.DeQueue())

		assert.Equal(t, 2, queue.Size())
	})

	t.Run("Test Queue IsEmpty ", func(t *testing.T) {
		queue := New[int]()
		assert.Equal(t, true, queue.IsEmpty())
		queue.EnQueue(1)
		queue.EnQueue(2)

		assert.Equal(t, false, queue.IsEmpty())
	})

	t.Run("Test Queue Length", func(t *testing.T) {
		queue := New[int]()
		queue.EnQueue(1)
		queue.EnQueue(2)
		queue.EnQueue(3)

		assert.Equal(t, 3, queue.Size())
	})
}
