package sort

import (
	"testing"

	"github.com/mucunga90/go/aligorithm"
)

func TestInsertionSort(t *testing.T) {
	execFunc := func(input []int) []int {
		return InsertionSort(input)
	}

	aligorithm.ExecutesSortTestCases(t, execFunc)
}
