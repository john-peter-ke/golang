package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximalSquare(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		input := [][]byte{
			{'1', '0', '1', '0', '0'},
			{'1', '0', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '0', '0', '1', '0'},
		}
		expected := 4
		assert.Equal(t, expected, maximalSquare(input))
	})

}
