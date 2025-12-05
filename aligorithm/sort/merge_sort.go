// worst-case time complexity: O(n log n)
// average-case time complexity: O(n log n)
// best-case time complexity: O(n log n)
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

	for k := 0; k < len(combined); k++ {
		if i >= len(left) {
			// Left array exhausted, copy remaining from right
			combined[k] = right[j]
			j++
		} else if j >= len(right) {
			// Right array exhausted, copy remaining from left
			combined[k] = left[i]
			i++
		} else if left[i] < right[j] {
			// Element in left array is smaller
			combined[k] = left[i]
			i++
		} else {
			// Element in right array is smaller or equal
			combined[k] = right[j]
			j++
		}
	}

	return combined
}
