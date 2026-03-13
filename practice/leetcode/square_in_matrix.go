package leetcode

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	rows, cols := len(matrix), len(matrix[0])
	dp := make([][]int, rows)
	maxSide := 0

	for i := range dp {
		dp[i] = make([]int, cols)
		for j := range dp[i] {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = minVal(dp[i-1][j], minVal(dp[i][j-1], dp[i-1][j-1])) + 1
				}

				if dp[i][j] > maxSide {
					maxSide = dp[i][j]
				}
			}
		}
	}

	return maxSide * maxSide
}

func minVal(a, b int) int {
	if a < b {
		return a
	}

	return b
}
