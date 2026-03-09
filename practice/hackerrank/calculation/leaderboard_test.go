package calculation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbingLeaderboard(t *testing.T) {

	t.Run("testing leaderboard", func(t *testing.T) {

		ranking := []int32{100, 100, 50, 40, 40, 20, 10}
		players := []int32{5, 25, 50, 120}

		expected := []int32{6, 4, 2, 1}
		assert.Equal(t, expected, climbingLeaderboard(ranking, players))
	})
}
