package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareArray(t *testing.T) {
	a := []int32{3, 2, 1}
	b := []int32{1, 2, 3}
	expectedArr := []int32{1, 1}

	t.Run("Compare Array", func(t *testing.T) {
		assert.Equal(t, expectedArr, CompareTriplets(a, b))
	})
}
