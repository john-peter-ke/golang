package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseList(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		expected := []int{5, 4, 3, 2, 1}
		list := reverseList(createList(arr))

		var count int
		current := list
		for current.Next != nil {
			assert.Equal(t, expected[count], current.Val)
			count++
			current = current.Next
		}
	})
}
