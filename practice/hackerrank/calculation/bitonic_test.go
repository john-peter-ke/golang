package calculation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPeakIndex(t *testing.T) {

	t.Run("FindPeakIndex", func(t *testing.T) {
		assert.Equal(t, int32(2), findPeakIndex([]int32{1, 3, 5, 4, 2}))
	})
}
