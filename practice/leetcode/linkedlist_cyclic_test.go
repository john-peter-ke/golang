package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedListHasCycle(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		arr := []int{3, 2, 0, -4}
		assert.Equal(t, true, linkedListHasCycle(createList(arr)))
	})

	t.Run("Case 2", func(t *testing.T) {
		arr := []int{3, 2, 0, -4}
		assert.Equal(t, true, linkedListHasCycleTwoPointer(createList(arr)))
	})

	t.Run("Case 3", func(t *testing.T) {
		arr := []int{1}
		assert.Equal(t, false, linkedListHasCycle(createList(arr)))
	})

}
