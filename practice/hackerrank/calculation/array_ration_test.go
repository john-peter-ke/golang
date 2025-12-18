package calculation

import (
	"testing"
)

func TestArrayRatio(t *testing.T) {
	input1 := []int32{-4, 3, -9, 0, 4, 1}

	t.Run("Array Ratio", func(t *testing.T) {
		PlusMinus(input1)
	})

}
