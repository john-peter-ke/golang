package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		s := "hello"
		expected := "olleh"
		assert.Equal(t, expected, reverseString(s))
	})

	t.Run("Case 2", func(t *testing.T) {
		s := "he"
		expected := "eh"
		assert.Equal(t, expected, reverseString(s))
	})

	t.Run("Case 3", func(t *testing.T) {
		s := "h"
		expected := "h"
		assert.Equal(t, expected, reverseString(s))
	})
}
