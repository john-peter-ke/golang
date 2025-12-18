package leetcode

func TwoSum(nums []int, target int) []int {
	diff := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		difference := target - nums[i]
		index, ok := diff[difference]
		if ok {
			return []int{index, i}
		}
		diff[nums[i]] = i
	}
	return []int{}
}

func TwoSumSorted(num []int, target int) []int {
	size := len(num)
	left := 0
	right := size - 1
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
