package _277_count_square_submatrices_with_all_ones

func countSquares(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	res := 0
	if m == 0 || n == 0 {
		return res
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = matrix[i][j]
			}
			if matrix[i][j] == 1 {
				if i >= 1 && j >= 1 {
					//动态规划
					dp[i][j] = min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
				}
			}
			res += dp[i][j]
		}
	}
	return res

}
func min(a, b, c int) int {
	min := a
	if min > b {
		min = b
	}
	if min > c {
		return c
	}
	return min
}
