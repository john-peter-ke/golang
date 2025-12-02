package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinearSearch(t *testing.T) {

	t.Run("LinearSearch", func(t *testing.T) {
		input := []int{8, 3, 1, 5, 4, 7, 2, 6, 9}
		expected := 5

		actual, err := LinearSearch(input, 7)
		assert.Equal(t, expected, actual)
		assert.NoError(t, err)
	})
}
