package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {

	t.Run("BinarySearch", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		expected := 4
		actual, err := BinarySearch(input, 5)
		assert.Equal(t, expected, actual)
		assert.NoError(t, err)
	})

	t.Run("BinarySearchResursion", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		expected := 4
		actual, err := BinarySearchResursion(input, 5)
		assert.Equal(t, expected, actual)
		assert.NoError(t, err)
	})
}
