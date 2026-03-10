package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Ensure nums1 is the shorter array
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	low, high := 0, m

	for low <= high {
		partition1 := (low + high) / 2
		partition2 := (m+n+1)/2 - partition1

		// Handle edge cases with -infinity and +infinity
		maxLeft1 := -1 << 63 // Smallest possible int
		if partition1 > 0 {
			maxLeft1 = nums1[partition1-1]
		}

		maxLeft2 := -1 << 63
		if partition2 > 0 {
			maxLeft2 = nums2[partition2-1]
		}

		minRight1 := 1<<63 - 1 // Largest possible int
		if partition1 < m {
			minRight1 = nums1[partition1]
		}

		minRight2 := 1<<63 - 1
		if partition2 < n {
			minRight2 = nums2[partition2]
		}

		// Check if we found the correct partition
		if maxLeft1 <= minRight2 && maxLeft2 <= minRight1 {
			if (m+n)%2 == 0 {
				// Even total length
				return float64(max(maxLeft1, maxLeft2)+min(minRight1, minRight2)) / 2.0
			} else {
				// Odd total length
				return float64(max(maxLeft1, maxLeft2))
			}
		} else if maxLeft1 > minRight2 {
			// Too far right in nums1, move left
			high = partition1 - 1
		} else {
			// Too far left in nums1, move right
			low = partition1 + 1
		}
	}

	return 0.0
}

// Helper functions for Go < 1.21 (Standard math package uses float64)
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
