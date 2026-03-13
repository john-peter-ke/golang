package leetcode

func TwoSum(nums []int, target int) []int {
	// Map to store: value -> index
	m := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		// Check if the complement already exists in the map
		if idx, found := m[complement]; found {
			return []int{idx, i}
		}

		// Store current number and its index
		m[num] = i
	}

	return nil
}

func TwoSumSorted(num []int, target int) []int {
	low, high := 0, len(num)-1

	for low < high {
		sum := num[low] + num[high]
		if sum == target {
			return []int{low, high}
		} else if sum > target {
			high--
		} else {
			low++
		}
	}

	return nil
}
