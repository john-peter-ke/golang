package calculation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountResponseTimeRegressions(t *testing.T) {

	t.Run("CountResponseTimeRegressions", func(t *testing.T) {
		assert.Equal(t, int32(2), countResponseTimeRegressions([]int32{100, 200, 150, 300}))
		assert.Equal(t, int32(1), countResponseTimeRegressions([]int32{1, 100}))
		assert.Equal(t, int32(0), countResponseTimeRegressions([]int32{0}))
		assert.Equal(t, int32(0), countResponseTimeRegressions([]int32{-1}))
		assert.Equal(t, int32(1), countResponseTimeRegressions([]int32{100, 200, 150, 0}))
		assert.Equal(t, int32(0), countResponseTimeRegressions([]int32{100, 100, 100, 100}))
		assert.Equal(t, int32(1), countResponseTimeRegressions([]int32{100, 0, 100, 0}))
	})
}
