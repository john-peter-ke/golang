package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		s := "anna"
		assert.Equal(t, true, isPalindrome(s))
	})

	t.Run("Case 2", func(t *testing.T) {
		s := "hello"
		assert.Equal(t, false, isPalindrome(s))
	})

	t.Run("Case 3", func(t *testing.T) {
		s := "race car"
		assert.Equal(t, true, isPalindrome(s))
	})
}
