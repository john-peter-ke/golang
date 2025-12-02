// worst-case time complexity: O(n)
// average-case time complexity: O(n)
// best-case time complexity: O(1)

// space complexity Iterative: O(1)

package search

import "github.com/mucunga90/go/constraint"

func LinearSearch[T constraint.Ordered](arr []T, query T) (int, error) {
	for i, item := range arr {
		if item == query {
			return i, nil
		}
	}

	return -1, ErrNotFound
}
