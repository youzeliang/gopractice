package _119_pascals_triangle_ii

func getRow(rowIndex int) []int {
	// 第0行
	a := []int{1}
	for i := 1; i <= rowIndex; i++ {
		// 尾部追加1
		a = append(a, 1)
		// 倒序计算杨辉三角当前行
		for j := i - 1; j > 0; j-- {
			a[j] += a[j-1]
		}
	}

	return a
}
