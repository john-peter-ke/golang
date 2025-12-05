package search

import (
	"testing"

	"github.com/mucunga90/go/aligorithm"
)

func TestLinearSearch(t *testing.T) {
	execFunc := func(input []int, searchItem int) (int, error) {
		return LinearSearch(input, searchItem)
	}

	aligorithm.ExecutesSearchTestCases(t, execFunc)
}
