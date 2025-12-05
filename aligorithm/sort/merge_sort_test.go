package sort

import (
	"testing"

	"github.com/mucunga90/go/aligorithm"
)

func TestMergeSort(t *testing.T) {
	execFunc := func(input []int) []int {
		return MergeSort(input)
	}

	aligorithm.ExecutesSortTestCases(t, execFunc)
}
