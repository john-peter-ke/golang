package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimeConversion(t *testing.T) {
	input := "07:05:45PM"
	expected := "19:05:45"

	t.Run("Time Conversion", func(t *testing.T) {
		assert.Equal(t, expected, TimeConversion(input))
	})
}
