package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoListsArray(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		arr1 := []int{1, 2, 3, 5, 6, 7}
		arr2 := []int{1, 3, 4, 6, 8, 9, 10}
		expected := []int{1, 1, 2, 3, 3, 4, 5, 6, 6, 7, 8, 9, 10}
		current := mergeTwoLists(createList(arr1), createList(arr2))

		var count int
		for current.Next != nil {
			assert.Equal(t, expected[count], current.Val)
			count++
			current = current.Next
		}
	})
}

func createList(arr []int) *ListNode {
	dummy := new(ListNode)
	result := dummy
	for i := len(arr) - 1; i >= 0; i-- {
		prev := result.Next
		result.Next = &ListNode{
			Val:  arr[i],
			Next: prev,
		}
	}

	return result.Next
}
