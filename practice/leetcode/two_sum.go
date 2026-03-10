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
	left, right := 0, len(num)-1

	for i := 0; i < right && right > left; i++ {
		if num[left]+num[right] > target {
			right = right - 1
		}

		if num[left]+num[right] == target {
			return []int{left, right}
		}
	}

	return nil
}
