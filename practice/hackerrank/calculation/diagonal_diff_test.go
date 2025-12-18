package calculation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiagonalDiff(t *testing.T) {
	input1 := [][]int32{
		{3},
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}
	var eDiff1 int32 = 15

	input2 := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
	}
	var eDiff2 int32 = 15

	input3 := [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12},
		{3},
	}
	var eDiff3 int32 = 15

	t.Run("Diagonal Diff", func(t *testing.T) {
		assert.Equal(t, eDiff1, DiagonalDifference(input1))
		assert.Equal(t, eDiff2, DiagonalDifference(input2))
		assert.Equal(t, eDiff3, DiagonalDifference(input3))
	})
}
