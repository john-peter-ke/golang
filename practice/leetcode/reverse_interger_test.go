package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseInterger(t *testing.T) {

	t.Run("Reverse Interger", func(t *testing.T) {
		assert.Equal(t, 321, reverse(123))
		assert.Equal(t, -321, reverse(-123))
		assert.Equal(t, 21, reverse(120))
	})
}
