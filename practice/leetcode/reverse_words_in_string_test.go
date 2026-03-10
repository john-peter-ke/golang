package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		s := "the sky is blue"
		expected := "blue is sky the"
		assert.Equal(t, expected, reverseWords(s))
	})

	t.Run("Case 2", func(t *testing.T) {
		s := "  hello world  "
		expected := "world hello"
		assert.Equal(t, expected, reverseWords(s))
	})
}
