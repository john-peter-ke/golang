package leetcode

import (
	"math"
)

func reverse(x int) int {
	reversed := 0
	for x != 0 {
		// Pop the last digit
		lastDigit := x % 10

		// Check for 32-bit overflow before performing the push
		// MaxInt32 = 2147483647, MinInt32 = -2147483648
		if reversed > math.MaxInt32/10 || (reversed == math.MaxInt32/10 && lastDigit > 7) {
			return 0
		}
		if reversed < math.MinInt32/10 || (reversed == math.MinInt32/10 && lastDigit < -8) {
			return 0
		}

		// Push the digit onto the result
		reversed = reversed*10 + lastDigit

		x /= 10
	}

	return reversed
}
