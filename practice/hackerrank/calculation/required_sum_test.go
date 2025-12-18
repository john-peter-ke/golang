package calculation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequiredSum(t *testing.T) {

	t.Run("RequiredSum", func(t *testing.T) {
		assert.Equal(t, 5, requiredSum([]int32{30, 31, 32, 33, 34, 35}))
	})
}
