package main

/*
 * Complete the 'birthdayCakeCandles' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY candles as parameter.
 */

func BirthdayCakeCandles(candles []int32) int32 {
	var tallest int32
	count := make(map[int32]int32)
	for _, c := range candles {
		if _, ok := count[c]; ok {
			count[c] = count[c] + 1

			if count[c] > tallest {
				tallest = count[c]
			}
		} else {
			count[c] = 1
		}
	}

	return tallest
}
