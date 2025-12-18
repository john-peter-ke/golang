package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {

	t.Run("TwoSum", func(t *testing.T) {

		input1 := []int{2, 7, 11, 15}
		target1 := 9
		expected1 := []int{0, 1}
		assert.Equal(t, TwoSum(input1, target1), expected1)

		input2 := []int{3, 2, 4}
		target2 := 6
		expected2 := []int{1, 2}
		assert.Equal(t, TwoSum(input2, target2), expected2)

		input3 := []int{3, 3}
		target3 := 6
		expected3 := []int{0, 1}
		assert.Equal(t, TwoSum(input3, target3), expected3)

		input4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		target4 := 8
		expected4 := []int{0, 6}
		assert.Equal(t, TwoSumSorted(input4, target4), expected4)

		input5 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		target5 := 0
		var expected5 []int
		assert.Equal(t, TwoSumSorted(input5, target5), expected5)
	})

}
