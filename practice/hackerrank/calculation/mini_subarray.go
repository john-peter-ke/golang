package calculation

import "math"

func minimumSubArray(nums []int32, K int32) int {
	// minLength stores the result and is initialized to infinity.
	minLength := math.MaxInt32
	found := false

	// Use a map to track element frequencies in the current window.
	counts := make(map[int32]int32)
	left := 0

	for right, num := range nums {
		// Expand the window to the right and update the count of the current element.
		counts[num]++

		// If the number of distinct elements exceeds K, shrink the window from the left.
		for int32(len(counts)) > K {
			leftNum := nums[left]
			counts[leftNum]--
			if counts[leftNum] == 0 {
				delete(counts, leftNum)
			}
			left++
		}

		// At this point, the window [left...right] has at most K distinct elements.
		// We now need to check if it has exactly K distinct elements.
		if int32(len(counts)) == K {
			found = true
			currentLength := right - left + 1
			if currentLength < minLength {
				minLength = currentLength
			}
		}
	}

	if !found {
		return -1 // No sub-array with exactly K distinct elements found.
	}
	return minLength
}

func minimumSubArray2(nums []int32, k int) int {
	count := make(map[int32]int)
	res, lFar, lNear := 0, 0, 0
	for r := 0; r < len(nums); r++ {
		count[nums[r]]++

		for len(count) > k {
			count[nums[lNear]]--
			if count[nums[lNear]] == 0 {
				delete(count, nums[lNear])
			}
			lNear++
			lFar = lNear
		}

		for count[nums[lNear]] > 1 {
			count[nums[lNear]]--
			lNear++
		}

		if len(count) == k {
			res += lNear - lFar + 1
		}
	}
	return res
}

func minimumSubArray3(nums []int32, k int) int {
	n := len(nums)

	count := 0

	for i := 0; i < n; i++ {
		mp := make(map[int32]int)
		for j := i; j < n; j++ {
			mp[nums[j]]++

			if len(mp) == k {
				count++
			}

			if len(mp) > k {
				break
			}
		}
	}

	return count
}
