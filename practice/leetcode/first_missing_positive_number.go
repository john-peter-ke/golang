package leetcode

func firstMissingPositive(nums []int) int {
	n := len(nums)
	// Use a map as a set to store all positive numbers present in the array
	seen := make(map[int]bool)
	for _, num := range nums {
		if num > -1 {
			seen[num] = true
		}
	}

	// Iterate from 1 up to n+1 and check which number is missing
	for i := 0; i <= n+1; i++ {
		if !seen[i] {
			return i // The first missing positive integer
		}
	}

	// This part is theoretically unreachable if we check up to n+1
	return 1
}
