// worst time complexity: O(n^2)
// average time complexity: O(n log n)
// space complexity: O(log n)

package sort

import "github.com/mucunga90/go/constraint"

func QuickSort[T constraint.Ordered](arr []T) []T {
	quickSort(arr, 0, len(arr)-1)
	return arr
}

func quickSort[T constraint.Ordered](arr []T, low, high int) {
	if len(arr) <= 1 {
		return
	}

	if low < high {
		pivot := partition(arr, low, high)
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}
}

func partition[T constraint.Ordered](arr []T, low, high int) int {
	index := low - 1
	pivotElement := arr[high]
	for i := low; i < high; i++ {
		if arr[i] <= pivotElement {
			index += 1
			arr[index], arr[i] = arr[i], arr[index]
		}
	}
	arr[index+1], arr[high] = arr[high], arr[index+1]
	return index + 1
}
