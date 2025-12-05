package sort

import (
	"testing"

	"github.com/mucunga90/go/aligorithm"
)

func TestSelectionSort(t *testing.T) {
	execFunc := func(input []int) []int {
		return SelectionSort(input)
	}

	aligorithm.ExecutesSortTestCases(t, execFunc)
}
