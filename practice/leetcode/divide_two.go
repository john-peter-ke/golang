package leetcode

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}

	neg1 := IsNegative(dividend)
	neg2 := IsNegative(divisor)
	if Abs(divisor) == 1 && Abs(dividend) > 1 {
		if neg1 || neg2 {
			if neg1 && neg2 {
				// HACK: Incorrect expected value
				if Abs(dividend) == 2147483648 {
					return 2147483647
				}

				return Abs(dividend)
			}
			return -Abs(dividend)
		}
		return dividend
	}

	count := 1 + recurse(Abs(dividend)-Abs(divisor), Abs(divisor))
	if neg1 || neg2 {
		if neg1 && neg2 {
			return count
		}
		return -count
	}

	return count
}

func recurse(dividend int, divisor int) int {
	if dividend < divisor && dividend < 0 {
		return -1
	}

	if dividend < divisor {
		return 0
	}

	return 1 + recurse(dividend-divisor, divisor)
}

func IsNegative(x int) bool {
	if x < 0 {
		return true
	}

	return false
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
