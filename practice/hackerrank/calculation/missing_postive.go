package calculation

func findSmallestMissingPositive(orderNumbers []int32) int32 {
	n := len(orderNumbers)
	// Use a map as a set to store all positive numbers present in the array
	seen := make(map[int32]bool)
	for _, num := range orderNumbers {
		if num > 0 {
			seen[num] = true
		}
	}

	// Iterate from 1 up to n+1 and check which number is missing
	var i int32 = 1
	for ; i <= int32(n+1); i++ {
		if !seen[i] {
			return i // The first missing positive integer
		}
	}

	// This part is theoretically unreachable if we check up to n+1
	return 1
}
