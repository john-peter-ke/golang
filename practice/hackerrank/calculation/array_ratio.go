package main

import (
	"fmt"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func PlusMinus(arr []int32) {
	size := len(arr)

	if size == 0 {
		printValues(0, 0, 0)
		return
	}

	var pos, neg, zero int
	for i := range arr {
		if arr[i] == 0 {
			zero++
		} else if arr[i] > 0 {
			pos++
		} else if arr[i] < 0 {
			neg++
		}
	}

	printValues(createRatio(pos, size), createRatio(neg, size), createRatio(zero, size))
}

func printValues(pos, neg, zero float64) {
	fmt.Printf("%s\n%s\n%s\n", fmt.Sprintf("%.6f", pos), fmt.Sprintf("%.6f", neg), fmt.Sprintf("%.6f", zero))
}

func createRatio(count, size int) float64 {
	if count == 0 {
		return 0
	}

	return float64(float64(count) / float64(size))
}
