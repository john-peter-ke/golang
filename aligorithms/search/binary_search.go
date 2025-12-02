// worst-case time complexity: O(n log n)
// average-case time complexity: O(n log n)

// space complexity Iterative: O(1)
// space complexity Recursive: O(log N)

package search

import "github.com/mucunga90/go/constraint"

func BinarySearch[T constraint.Ordered](arr []T, query T) (int, error) {
	startIndex, endIndex := 0, len(arr)-1
	var mid int
	for startIndex <= endIndex {
		mid = int(startIndex + (endIndex-startIndex)/2)
		if arr[mid] > query {
			endIndex = mid - 1
		} else if arr[mid] < query {
			startIndex = mid + 1
		} else {
			return mid, nil
		}
	}
	return -1, ErrNotFound
}

func BinarySearchResursion[T constraint.Ordered](arr []T, query T) (int, error) {
	return Binary(arr, query, 0, len(arr)-1)
}

func Binary[T constraint.Ordered](array []T, target T, lowIndex int, highIndex int) (int, error) {
	if highIndex < lowIndex || len(array) == 0 {
		return -1, ErrNotFound
	}
	mid := int(lowIndex + (highIndex-lowIndex)/2)
	if array[mid] > target {
		return Binary(array, target, lowIndex, mid-1)
	} else if array[mid] < target {
		return Binary(array, target, mid+1, highIndex)
	} else {
		return mid, nil
	}
}
