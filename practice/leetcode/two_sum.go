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
