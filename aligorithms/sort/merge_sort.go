// worst-case time complexity: O(n^2)
// average-case time complexity: O(n^2)
// best-case time complexity: O(n^2)
// space complexity: O(n)

package sort

import "github.com/mucunga90/go/constraint"

func MergeSort[T constraint.Ordered](arr []T) []T {

	if len(arr) < 2 {
		return arr
	}

	middle := len(arr) / 2
	left := MergeSort(arr[:middle])
	right := MergeSort(arr[middle:])

	return merge(left, right)
}

func merge[T constraint.Ordered](left []T, right []T) []T {
	combined := make([]T, len(left)+len(right))
	var i, j = 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			combined[i+j] = left[i]
			i++
		} else {
			combined[i+j] = right[j]
			j++
		}
	}

	for i < len(left) {
		combined[i+j] = left[i]
		i++
	}

	for j < len(right) {
		combined[i+j] = right[j]
		j++
	}

	return combined
}
