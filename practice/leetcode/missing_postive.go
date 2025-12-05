package leetcode

import "sort"

func firstMissingPositive(nums []int) int {
	sort.Ints(nums)
	if nums[0] > 1 {
		return 1
	}

	if len(nums) == 1 && nums[0] == 0 {
		return 1
	}

	if len(nums) == 1 && nums[0] < 0 {
		return 1
	}

	for i := 0; i < len(nums); i++ {
		if i+1 < len(nums) && nums[i+1] > 0 && nums[i] > 0 && (nums[i+1]-nums[i]) > 1 {
			return nums[i] + 1
		}
	}

	if len(nums) == 1 {
		return len(nums) + 1
	}

	return len(nums)
}
