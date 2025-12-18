package calculation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBirthdayCakeCandles(t *testing.T) {
	input1 := []int32{3, 2, 1, 3}
	var exp1 int32 = 2

	t.Run("Birthday Cake Candles", func(t *testing.T) {
		assert.Equal(t, exp1, BirthdayCakeCandles(input1))
	})
}
