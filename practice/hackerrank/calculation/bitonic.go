package calculation

func findPeakIndex(counts []int32) int32 {
	size := int32(len(counts))
	var low, high int32 = 0, size - 1
	for low < high {
		mid := low + (high-low)/2

		if counts[mid] < counts[mid+1] {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}
