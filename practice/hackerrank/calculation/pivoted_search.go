package calculation

func searchRotatedTimestamps(nums []int32, target int32) int32 {
	low, high := 0, len(nums)-1
	var mid int
	for low <= high {
		mid = low + (high-low)/2
		if nums[mid] == target {
			return int32(mid)
		}

		if nums[low] <= nums[mid] {
			if target >= nums[low] && nums[mid] > target {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	return -1
}
