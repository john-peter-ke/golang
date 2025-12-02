package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleton(t *testing.T) {

	t.Run("Singleton", func(t *testing.T) {
		first := GetInstance()
		second := GetInstance()
		assert.Equal(t, first, second)
	})
}
