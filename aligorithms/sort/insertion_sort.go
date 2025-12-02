// worst-case time complexity: O(n^2)
// average-case time complexity: O(n^2)
// best-case time complexity: O(n)
// space complexity: O(1)

package sort

import "github.com/mucunga90/go/constraint"

func InsertionSort[T constraint.Ordered](arr []T) []T {
	for i := 1; i < len(arr); i++ {
		value := arr[i]
		index := i
		for ; index > 0 && arr[index-1] > value; index-- {
			arr[index] = arr[index-1]
		}
		arr[index] = value
	}
	return arr
}
