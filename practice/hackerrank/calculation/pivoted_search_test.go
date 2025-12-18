package calculation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchRotatedTimestamps(t *testing.T) {

	t.Run("SearchRotatedTimestamps", func(t *testing.T) {
		assert.Equal(t, int32(3), searchRotatedTimestamps([]int32{1609466400, 1609470000, 1609473600, 1609459200, 1609462800}, 1609459200))
	})
}
