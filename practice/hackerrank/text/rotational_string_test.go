package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNonTrivialRotation(t *testing.T) {

	t.Run("Missing Positive", func(t *testing.T) {
		assert.Equal(t, false, isNonTrivialRotation("a", "a"))
		assert.Equal(t, false, isNonTrivialRotation("a", "b"))
	})
}
