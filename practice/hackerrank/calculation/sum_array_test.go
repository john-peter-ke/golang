package calculation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumArray(t *testing.T) {
	input := []int32{1, 2, 3, 4, 5}
	var expectedSum int32 = 15

	t.Run("Sum array", func(t *testing.T) {
		assert.Equal(t, expectedSum, SimpleArraySum(input))
	})
}
