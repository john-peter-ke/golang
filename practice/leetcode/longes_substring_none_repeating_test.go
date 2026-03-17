package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		input := "pwwkew"
		expected := 3
		assert.Equal(t, lengthOfLongestSubstringWindow(input), expected)
		assert.Equal(t, lengthOfLongestSubstringOptimized(input), expected)
	})

	t.Run("Case 2", func(t *testing.T) {

		input := "au"
		expected := 2
		assert.Equal(t, lengthOfLongestSubstringWindow(input), expected)
		assert.Equal(t, lengthOfLongestSubstringOptimized(input), expected)
		assert.Equal(t, expected, lengthOfLongestSubstringHashTable(input))
	})
}
