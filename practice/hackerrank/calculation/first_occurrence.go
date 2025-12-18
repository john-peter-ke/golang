package calculation

func findFirstOccurrence(nums []int32, target int32) int32 {
	startIndex, endIndex := 0, len(nums)-1
	var midIndex int
	for startIndex <= endIndex {
		midIndex = startIndex + (endIndex-startIndex)/2
		if nums[midIndex] > target {
			endIndex = midIndex - 1
		} else if nums[midIndex] < target {
			startIndex = midIndex + 1
		} else if nums[midIndex] == nums[startIndex] {
			return int32(startIndex)
		} else {
			return int32(midIndex)
		}
	}

	return -1
}
