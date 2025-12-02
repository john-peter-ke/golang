package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {

	t.Run("Test IsEmpty", func(t *testing.T) {
		stack := NewStack[int]()

		assert.Equal(t, true, stack.IsEmpty())
	})

	t.Run("Test Push", func(t *testing.T) {
		stack := NewStack[int]()

		assert.Equal(t, 0, stack.Size())

		stack.Push(1)
		stack.Push(2)
		stack.Push(3)

		assert.Equal(t, 3, stack.Size())
	})

	t.Run("Test Peek", func(t *testing.T) {
		stack := NewStack[string]()
		assert.Equal(t, "", stack.Peek())

		stack.Push("one")
		assert.Equal(t, "one", stack.Peek())

		stack.Push("two")
		assert.Equal(t, "two", stack.Peek())

		stack.Push("three")
		assert.Equal(t, "three", stack.Peek())
	})

	t.Run("Test Pop", func(t *testing.T) {
		stack := NewStack[int]()
		stack.Push(1)
		stack.Push(2)

		assert.Equal(t, 2, stack.Pop())
		assert.Equal(t, 1, stack.Pop())
		assert.Equal(t, 0, stack.Pop())
	})
}
