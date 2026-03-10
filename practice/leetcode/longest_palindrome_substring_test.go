package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindromeSubstring(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		s := "cbbd"
		expected := "bb"
		assert.Equal(t, expected, longestPalindromeSubstring(s))
	})

	t.Run("Case 2", func(t *testing.T) {
		s := "babad"
		expected := "bab"
		assert.Equal(t, expected, longestPalindromeSubstring(s))
	})

}
