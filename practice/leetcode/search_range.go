package leetcode

func searchRange(nums []int, target int) []int {
	start, end := 0, len(nums)-1
	var mid int
	first := -1
	for start <= end {
		mid = int(start + (end-start)/2)
		if nums[mid] == target {
			first = mid
			end = mid - 1
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	start, end = 0, len(nums)-1
	last := -1
	for start <= end {
		mid = int(start + (end-start)/2)
		if nums[mid] == target {
			last = mid
			start = mid + 1
		} else if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return []int{first, last}
}
