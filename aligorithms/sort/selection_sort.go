// worst-case time complexity: O(n^2)
// average-case time complexity: O(n^2)
// best-case time complexity: O(n)
// space complexity: O(1)

package sort

import (
	"github.com/mucunga90/go/constraint"
)

func SelectionSort[T constraint.Ordered](arr []T) []T {
	for i := range arr {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		arr[i], arr[min] = arr[min], arr[i]
	}

	return arr
}
