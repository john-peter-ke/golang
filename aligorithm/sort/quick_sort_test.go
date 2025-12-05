package sort

import (
	"testing"

	"github.com/mucunga90/go/aligorithm"
)

func TestQuickSort(t *testing.T) {
	execFunc := func(input []int) []int {
		return QuickSort(input)
	}

	aligorithm.ExecutesSortTestCases(t, execFunc)
}
