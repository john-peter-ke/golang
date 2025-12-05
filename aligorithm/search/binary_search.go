// worst-case time complexity: O(n log n)
// average-case time complexity: O(n log n)

// space complexity Iterative: O(1)
// space complexity Recursive: O(log N)

package search

import (
	"github.com/mucunga90/go/aligorithm"
	"github.com/mucunga90/go/constraint"
)

func BinarySearch[T constraint.Ordered](arr []T, target T) (int, error) {
	startIndex, endIndex := 0, len(arr)-1
	var mid int
	for startIndex <= endIndex {
		mid = int(startIndex + (endIndex-startIndex)/2)
		if arr[mid] > target {
			endIndex = mid - 1
		} else if arr[mid] < target {
			startIndex = mid + 1
		} else {
			return mid, nil
		}
	}
	return -1, aligorithm.ErrNotFound
}
