package calculation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSubarraysWithSumAndMaxAtMost(t *testing.T) {
	t.Run("CountSubarraysWithSumAndMaxAtMost", func(t *testing.T) {
		assert.Equal(t, int64(2), countSubarraysWithSumAndMaxAtMost([]int32{2, -1, 2, 1, -2, 3}, int64(3), int64(2)))
	})
}
