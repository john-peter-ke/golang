package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	var input1 uint32 = 1
	var input2 uint32 = 3
	expectedSum := input1 + input2

	t.Run("Sum", func(t *testing.T) {
		assert.Equal(t, expectedSum, sumNumbers(input1, input2))
	})
}
