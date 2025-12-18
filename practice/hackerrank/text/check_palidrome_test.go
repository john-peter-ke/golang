package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAlphabeticPalindrome(t *testing.T) {

	t.Run("Missing Positive", func(t *testing.T) {
		assert.Equal(t, true, isAlphabeticPalindrome("abc123cba"))
	})
}
