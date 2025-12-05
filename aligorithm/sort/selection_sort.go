package sort

import (
	"github.com/mucunga90/go/constraint"
)

// For an array of n/2 values, there are n comparisons: n/2.n
// worst-case time complexity: O(n^2)
// average-case time complexity: O(n^2)
// best-case time complexity: O(n^2)
// space complexity: O(1)
// InsertionSort received unsorted arr and returns sorted arr
func SelectionSort[T constraint.Ordered](arr []T) []T {
	for i := range arr {
		min := i

		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j // goal is to identify the minimum value index
			}
		}

		arr[i], arr[min] = arr[min], arr[i]
	}

	return arr
}
