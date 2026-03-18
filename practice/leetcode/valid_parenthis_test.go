package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidParenthesis(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		input := "()"
		expected := true
		assert.Equal(t, expected, isValidParenthesis(input))
	})

	t.Run("Case 2", func(t *testing.T) {
		input := "()[]{}"
		expected := true
		assert.Equal(t, expected, isValidParenthesis(input))
	})

	t.Run("Case 3", func(t *testing.T) {
		input := "(]"
		expected := false
		assert.Equal(t, expected, isValidParenthesis(input))
	})

	t.Run("Case 4", func(t *testing.T) {
		input := "([])"
		expected := true
		assert.Equal(t, expected, isValidParenthesis(input))
	})

	t.Run("Case 5", func(t *testing.T) {
		input := "([)]"
		expected := false
		assert.Equal(t, expected, isValidParenthesis(input))
	})

	t.Run("Case 6", func(t *testing.T) {
		input := "["
		expected := false
		assert.Equal(t, expected, isValidParenthesis(input))
	})
}
