package sort

import (
	"testing"

	"github.com/mucunga90/go/aligorithm"
)

func TestHeapSort(t *testing.T) {
	execFunc := func(input []int) []int {
		return HeapSort(input)
	}

	aligorithm.ExecutesSortTestCases(t, execFunc)
}
