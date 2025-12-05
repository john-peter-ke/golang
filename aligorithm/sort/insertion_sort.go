package sort

import "github.com/mucunga90/go/constraint"

// For an array of n/2 values, there are n comparisons: n/2.n
// worst-case time complexity: O(n^2)
// average-case time complexity: O(n^2)
// best-case time complexity: O(n)
// space complexity: O(1)
// InsertionSort received unsorted arr and returns sorted arr
func InsertionSort[T constraint.Ordered](arr []T) []T {
	for i := range arr {
		index := i
		value := arr[i]

		for ; index > 0 && arr[index-1] > value; index-- {
			arr[index] = arr[index-1]
		}

		arr[index] = value
	}

	return arr
}
