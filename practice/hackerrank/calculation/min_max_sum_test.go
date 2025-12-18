package calculation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinMaxSum(t *testing.T) {
	input1 := []int32{1, 2, 3, 4, 5}
	var eMin int64 = 10
	var eMax int64 = 14

	t.Run("Min Max Sum", func(t *testing.T) {
		min, max := MiniMaxSum(input1)
		assert.Equal(t, eMin, min)
		assert.Equal(t, eMax, max)
	})
}
