package calculation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckGame(t *testing.T) {
	t.Run("Check game", func(t *testing.T) {
		assert.Equal(t, "Second", ChessboardGame(5, 2))
		assert.Equal(t, "First", ChessboardGame(5, 3))
		assert.Equal(t, "First", ChessboardGame(15, 15))
	})
}
