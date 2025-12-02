package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {

	input := []int{4, 6, 2, 9, 3, 1, 5, 7, 8}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	t.Run("Insertion Sort", func(t *testing.T) {
		assert.Equal(t, expected, InsertionSort(input))
	})
}
