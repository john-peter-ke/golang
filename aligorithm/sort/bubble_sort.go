package sort

import "github.com/mucunga90/go/constraint"

// For an array of n values, there are n comparisons: n.n
// worst-case time complexity: O(n^2)
// average-case time complexity: O(n^2)
// best-case time complexity: O(n) e.g [7,3,9,12,11] sorted after first run
// space complexity: O(1)
// BubbleSort received unsorted arr and returns sorted arr
func BubbleSort[In constraint.Ordered](arr []In) []In {
	swapped := true
	for swapped {
		swapped = false // arr is considered sorted whenever no swap happened
		// the comparison arr is one item to avoid out of bound
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
	}

	return arr
}
