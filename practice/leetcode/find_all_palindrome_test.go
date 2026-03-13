package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAllPalindrome(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		input := "babad"
		assert.Equal(t, 3, len(findAllPalindrome(input)))
	})
}
