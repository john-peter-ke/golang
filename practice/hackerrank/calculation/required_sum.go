package calculation

import (
	"strconv"
	"strings"
)

func requiredSum(num []int32) int32 {
	count := make(map[int32]int32)
	for i := range num {
		if _, ok := count[num[i]]; !ok {
			countDigits(count, num[i])
		}
	}

	var sum int32
	for _, val := range count {
		sum += val
	}

	return sum
}

func countDigits(count map[int32]int32, num int32) {
	for i := range 5000 {
		digits := strings.Split(strconv.Itoa(int(i)), "")

		var sum int32
		for _, value := range digits {
			digit, _ := strconv.ParseInt(value, 10, 32)
			sum += int32(digit)
		}

		if sum == num {
			count[num] += 1
		}
	}
}
