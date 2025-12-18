package calculation

func countSubarraysWithSumAndMaxAtMost(nums []int32, k int64, M int64) int64 {
	count := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		var currentSum int64 = 0
		var currentMax int64 = -1
		for j := i; j < n; j++ {
			currentSum += int64(nums[j])
			if int64(nums[j]) > int64(currentMax) {
				currentMax = int64(nums[j])
			}
			// Check conditions: sum is k AND max element is exactly targetMax
			if currentSum == k && currentMax == int64(M) {
				count++
			}

			if currentMax > int64(M) {
				break
			}
		}
	}
	return int64(count)
}
