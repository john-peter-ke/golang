package calculation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindFirstOccurrence(t *testing.T) {

	t.Run("FindFirstOccurrence", func(t *testing.T) {
		assert.Equal(t, int32(-1), findFirstOccurrence([]int32{0, 5}, 1))
	})
}
