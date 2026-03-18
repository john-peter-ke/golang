package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoLinkedList(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		arr1 := []int{2, 4, 3}
		arr2 := []int{5, 6, 4}
		expected := []int{7, 0, 8}

		result := addTwoLinkedList(createList(arr1), createList(arr2))

		var count int
		for result.Next != nil {
			assert.Equal(t, expected[count], result.Val)
			count++
			result = result.Next
		}
	})
}
