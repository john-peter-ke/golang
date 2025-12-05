package search

import (
	"testing"

	"github.com/mucunga90/go/aligorithm"
)

func TestBinarySearch(t *testing.T) {
	execFunc := func(input []int, searchItem int) (int, error) {
		return BinarySearch(input, searchItem)
	}
	aligorithm.ExecutesSearchTestCases(t, execFunc)
}
