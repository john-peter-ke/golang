package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstMissingPositive(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		arr := []int{1, 2, 0}
		expected := 3
		assert.Equal(t, expected, firstMissingPositive(arr))
	})

	t.Run("Case 2", func(t *testing.T) {
		arr := []int{3, 4, -1, 1}
		expected := 2
		assert.Equal(t, expected, firstMissingPositive(arr))
	})

	t.Run("Case 3", func(t *testing.T) {
		arr := []int{7, 8, 9, 11, 12}
		expected := 1
		assert.Equal(t, expected, firstMissingPositive(arr))
	})
}
