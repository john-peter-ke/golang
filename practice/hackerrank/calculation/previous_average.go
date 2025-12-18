package calculation

func countResponseTimeRegressions(responseTimes []int32) int32 {
	if len(responseTimes) <= 1 {
		return 0
	}

	var count int32 = 0
	var sum float64 = float64(responseTimes[0])

	for i := 1; i < len(responseTimes); i++ {
		if float64(responseTimes[i]) > sum/float64(i) {
			count++
		}

		sum += float64(responseTimes[i])
	}

	return count
}
