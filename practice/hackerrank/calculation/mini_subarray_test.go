package calculation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumSubArray(t *testing.T) {

	t.Run("RequiredSum", func(t *testing.T) {
		assert.Equal(t, 2, minimumSubArray3([]int32{1, 2, 2, 3}, 2))
	})
}
