package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapSort(t *testing.T) {

	t.Run("MergeSort", func(t *testing.T) {
		input := []int{8, 3, 1, 5, 4, 7, 2, 6, 9}
		expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		assert.Equal(t, expected, HeapSort(input))
	})
}
