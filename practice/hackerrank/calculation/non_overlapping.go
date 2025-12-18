package calculation

func maximizeNonOverlappingMeetings(meetings [][]int32) int32 {
	quickSort(meetings, 0, len(meetings)-1)
	var count, pvend int32
	for i := 0; i < len(meetings); i++ {
		if len(meetings[i]) <= 1 {
			//skip
			continue
		}

		if meetings[i][1] > meetings[i][0] && meetings[i][0] >= pvend {
			pvend = meetings[i][1]
			count++
		}
	}

	return count
}

func quickSort(mt [][]int32, low, high int) {
	if len(mt) <= 1 {
		return
	}

	if high > low {
		pivot := partition(mt, low, high)
		quickSort(mt, low, pivot-1)
		quickSort(mt, pivot+1, high)
	}
}

func partition(mt [][]int32, low, high int) int {
	index := low - 1
	pivotElement := mt[high]
	for i := low; i < high; i++ {
		if len(mt[i]) == 1 {
			mt[i] = []int32{mt[i][0], 0}
		}
		if len(pivotElement) == 1 {
			pivotElement = []int32{pivotElement[0], 0}
		}
		if mt[i][1] <= pivotElement[1] {
			index += 1
			mt[index], mt[i] = mt[i], mt[index]
		}
	}

	mt[index+1], mt[high] = mt[high], mt[index+1]
	return index + 1
}
