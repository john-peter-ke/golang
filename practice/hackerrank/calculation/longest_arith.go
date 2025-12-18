package calculation

func findLongestArithmeticProgression(arr []int32, k int32) int32 {
	if len(arr) == 0 {
		return 0
	}

	dp := make(map[int32]int32)
	var maxLength int32

	arrSort(arr, 0, len(arr)-1)
	for _, num := range arr {
		var length int32 = 1
		if val, ok := dp[num-k]; ok {
			length = val + 1
		}

		dp[num] = length
		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength

}

func arrSort(mt []int32, low, high int) {
	if len(mt) <= 1 {
		return
	}

	if high > low {
		pivot := arrPartition(mt, low, high)
		arrSort(mt, low, pivot-1)
		arrSort(mt, pivot+1, high)
	}
}

func arrPartition(mt []int32, low, high int) int {
	index := low - 1
	pivotElement := mt[high]
	for i := low; i < high; i++ {
		if mt[i] <= pivotElement {
			index += 1
			mt[index], mt[i] = mt[i], mt[index]
		}
	}

	mt[index+1], mt[high] = mt[high], mt[index+1]
	return index + 1
}
