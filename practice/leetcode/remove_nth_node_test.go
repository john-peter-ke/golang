package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveNthFromEnd(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		arr := []int{1, 2, 3, 5, 6, 7}
		expected := []int{1, 2, 3, 5, 7}
		current := removeNthFromEnd(createList(arr), 2)

		var count int
		for current.Next != nil {
			assert.Equal(t, expected[count], current.Val)
			count++
			current = current.Next
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		arr := []int{1, 2, 3, 5, 6, 7}
		expected := []int{1, 2, 3, 5, 7}
		current := removeNthFromEndTwoPointer(createList(arr), 2)

		var count int
		for current.Next != nil {
			assert.Equal(t, expected[count], current.Val)
			count++
			current = current.Next
		}
	})
}
