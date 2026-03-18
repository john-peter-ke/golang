package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstMissingPositive(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		arr := []int{0, 1}
		expected := 2
		assert.Equal(t, expected, firstMissingPositive(arr))
	})

	t.Run("Case 2", func(t *testing.T) {
		arr := []int{3, 0, 1}
		expected := 2
		assert.Equal(t, expected, firstMissingPositive(arr))
	})

	t.Run("Case 3", func(t *testing.T) {
		arr := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
		expected := 8
		assert.Equal(t, expected, firstMissingPositive(arr))
	})

	t.Run("Case 4", func(t *testing.T) {
		arr := []int{1}
		expected := 0
		assert.Equal(t, expected, firstMissingPositive(arr))
	})
}
