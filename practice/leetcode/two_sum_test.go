package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {

		input1 := []int{2, 7, 11, 15}
		target1 := 9
		expected1 := []int{0, 1}
		assert.Equal(t, expected1, TwoSum(input1, target1))

		input2 := []int{3, 2, 4}
		target2 := 6
		expected2 := []int{1, 2}
		assert.Equal(t, expected2, TwoSum(input2, target2))

		input3 := []int{3, 3}
		target3 := 6
		expected3 := []int{0, 1}
		assert.Equal(t, expected3, TwoSum(input3, target3))

		input4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		target4 := 8
		expected4 := []int{0, 6}
		assert.Equal(t, expected4, TwoSumSorted(input4, target4))

		input5 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		target5 := 0
		var expected5 []int
		assert.Equal(t, expected5, TwoSumSorted(input5, target5))
	})

	t.Run("Case 2", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		target := 24
		var expected []int
		assert.Equal(t, expected, TwoSumSorted(input, target))
	})

}
