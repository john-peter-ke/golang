package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMedianTwoSortedArray(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		num1 := []int{1, 3}
		num2 := []int{2}
		expected1 := float64(2)
		assert.Equal(t, expected1, findMedianSortedArrays(num1, num2))
	})

	t.Run("Case 2", func(t *testing.T) {
		num1 := []int{1, 2}
		num2 := []int{3, 4}
		expected1 := float64(2.5)
		assert.Equal(t, expected1, findMedianSortedArrays(num1, num2))
	})

	t.Run("Case 3", func(t *testing.T) {
		num1 := []int{0, 0, 0, 0, 0}
		num2 := []int{-1, 0, 0, 0, 0, 0, 1}
		expected1 := float64(0)
		assert.Equal(t, expected1, findMedianSortedArrays(num1, num2))
	})
}
