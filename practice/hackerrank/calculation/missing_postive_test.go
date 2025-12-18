package calculation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingPositive(t *testing.T) {

	t.Run("Missing Positive", func(t *testing.T) {
		assert.Equal(t, int32(4), findSmallestMissingPositive([]int32{-2, -1, 1, 2, 3}))
		assert.Equal(t, int32(3), findSmallestMissingPositive([]int32{1, 2, 0}))
		assert.Equal(t, int32(2), findSmallestMissingPositive([]int32{3, 4, -2, 1}))
		assert.Equal(t, int32(2), findSmallestMissingPositive([]int32{3, 4, -1, 1}))
		assert.Equal(t, int32(1), findSmallestMissingPositive([]int32{3, 4, -2, -1}))
		assert.Equal(t, int32(1), findSmallestMissingPositive([]int32{7, 8, 9, 11, 12}))
		assert.Equal(t, int32(2), findSmallestMissingPositive([]int32{1}))
		assert.Equal(t, int32(2), findSmallestMissingPositive([]int32{1, 1}))
		assert.Equal(t, int32(1), findSmallestMissingPositive([]int32{0}))
		assert.Equal(t, int32(1), findSmallestMissingPositive([]int32{0, 0}))
		assert.Equal(t, int32(4), findSmallestMissingPositive([]int32{100000, 3, 4000, 2, 15, 1, 99999}))
	})
}
