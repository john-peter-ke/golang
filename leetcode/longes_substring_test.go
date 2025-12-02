package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {

	t.Run("TwoSum", func(t *testing.T) {

		input1 := "pwwkew"
		expected1 := 3
		assert.Equal(t, lengthOfLongestSubstring(input1), expected1)

		input2 := "au"
		expected2 := 2
		assert.Equal(t, lengthOfLongestSubstring(input2), expected2)
	})

}
