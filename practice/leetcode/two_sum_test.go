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
	})

}
