package calculation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximizeNonOverlappingMeetings(t *testing.T) {

	t.Run("MaximizeNonOverlappingMeetings", func(t *testing.T) {

		assert.Equal(t, int32(3), maximizeNonOverlappingMeetings([][]int32{
			{3},
			{2},
			{1, 2},
			{2, 3},
			{3, 4},
		}))

		assert.Equal(t, int32(1), maximizeNonOverlappingMeetings([][]int32{
			{1},
			{2},
			{5, 10},
		}))

		assert.Equal(t, int32(2), maximizeNonOverlappingMeetings([][]int32{
			{1, 100},
			{11, 22},
			{1, 11},
			{2, 12},
		}))
	})
}
