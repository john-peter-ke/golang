package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackList(t *testing.T) {

	t.Run("Test IsEmpty", func(t *testing.T) {
		stack := NewStackList()

		assert.Equal(t, true, stack.IsEmpty())
	})

	t.Run("Test Push", func(t *testing.T) {
		stack := NewStackList()

		stack.Push("one")
		stack.Push("two")

		assert.Equal(t, 2, stack.Size())
	})

	t.Run("Test Peek", func(t *testing.T) {
		stack := NewStackList()

		stack.Push("1")
		assert.Equal(t, "1", stack.Peek())

		stack.Push("2")
		assert.Equal(t, "2", stack.Peek())

		stack.Pop()
		stack.Pop()
		assert.Equal(t, nil, stack.Peek())
	})

	t.Run("Test Pop", func(t *testing.T) {
		stack := NewStackList()

		stack.Push("1")
		stack.Push("2")

		assert.Equal(t, "2", stack.Pop())
		assert.Equal(t, "1", stack.Pop())
		assert.Equal(t, nil, stack.Pop())
	})
}
