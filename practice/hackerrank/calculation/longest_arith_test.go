package calculation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLongestArithmeticProgression(t *testing.T) {

	t.Run("FindLongestArithmeticProgression", func(t *testing.T) {
		assert.Equal(t, int32(6), findLongestArithmeticProgression([]int32{8, 1, -1, 0, 3, 6, 2, 4, 5, 7, 9}, 2))
		//assert.Equal(t, int32(1), findLongestArithmeticProgression([]int32{1, 42, 7}, 2))
	})
}
