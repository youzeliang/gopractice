package _120_triangle

func minimumTotal(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- { // 倒数二行开始
		for j := 0; j <= i; j++ {
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}

func min(x, y int) int {
	if x > y {
		x = y
	}
	return x
}
