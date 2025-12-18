package calculation

import (
	"fmt"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func MiniMaxSum(arr []int32) (int64, int64) {
	newArr := insertionSort(arr)

	var totalSum int64
	for _, num := range newArr {
		totalSum = totalSum + int64(num)
	}
	minSum := totalSum - int64(newArr[len(newArr)-1])
	maxSum := totalSum - int64(newArr[0])

	fmt.Printf("%d %d", minSum, maxSum)
	return minSum, maxSum
}

func insertionSort(arr []int32) []int32 {
	for i := 1; i < len(arr); i++ {
		value := arr[i]
		index := i
		for ; index > 0 && arr[index-1] > value; index-- {
			arr[index] = arr[index-1]
		}
		arr[index] = value
	}
	return arr
}
