package calculation

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func DiagonalDifference(arr [][]int32) int32 {
	rows := len(arr)
	if rows == 0 {
		return 0
	}

	var newArr [][]int32
	var squareSize int
	for i := range arr {
		if len(arr[i]) < 2 {
			continue
		}

		if squareSize == 0 {
			squareSize = len(arr[i])
		}

		if len(newArr) < squareSize {
			newArr = append(newArr, arr[i])
		}
	}

	var lsum, rsum int32
	for i, j := 0, len(newArr)-1; i < len(newArr); i, j = i+1, j-1 {
		lsum += newArr[i][i]
		rsum += newArr[i][j]
	}

	return abs(lsum - rsum)
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}

	return x
}
