package sort

import (
	"testing"

	"github.com/mucunga90/go/aligorithm"
)

func TestBubbleSort(t *testing.T) {
	execFunc := func(input []int) []int {
		return BubbleSort(input)
	}

	aligorithm.ExecutesSortTestCases(t, execFunc)
}
