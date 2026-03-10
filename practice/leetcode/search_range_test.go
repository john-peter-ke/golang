package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchRange(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		input := []int{5, 7, 7, 8, 8, 10}
		target := 8
		expected := []int{3, 4}
		assert.Equal(t, expected, searchRange(input, target))
	})

}
