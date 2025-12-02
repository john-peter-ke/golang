package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackNode(t *testing.T) {

	t.Run("Test IsEmpty", func(t *testing.T) {
		stack := NewStackNode[int]()

		assert.Equal(t, true, stack.IsEmpty())

		stack.Push(1)

		assert.Equal(t, false, stack.IsEmpty())
	})

	t.Run("Test Push", func(t *testing.T) {
		stack := NewStackNode[string]()

		assert.Equal(t, 0, stack.Size())

		stack.Push("one")
		stack.Push("two")

		assert.Equal(t, 2, stack.Size())
	})

	t.Run("Test Peek", func(t *testing.T) {
		stack := NewStackNode[int]()

		stack.Push(1)

		assert.Equal(t, 1, stack.Peek())

		stack.Push(2)

		assert.Equal(t, 2, stack.Peek())
	})

	t.Run("Test Pop", func(t *testing.T) {
		stack := NewStackNode[string]()

		stack.Push("10")
		stack.Push("20")
		stack.Push("30")
		stack.Push("40")
		stack.Push("50")

		assert.Equal(t, "50", stack.Pop())
		assert.Equal(t, "40", stack.Peek())

		assert.Equal(t, "40", stack.Pop())
		assert.Equal(t, "30", stack.Peek())
	})

	t.Run("Test Print", func(t *testing.T) {
		stack := NewStackNode[string]()
		stack.Push("10")
		stack.Push("20")
		stack.Push("30")
		stack.Push("40")
		stack.Push("50")

		stack.Print()
	})
}
