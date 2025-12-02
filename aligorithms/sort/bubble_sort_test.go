package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	input := []int{1, 4, 5, 7, 2, 9, 3, 6, 8}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	t.Run("Bubble sort", func(t *testing.T) {
		assert.Equal(t, expected, BubbleSort(input))
	})
}
