package _931_minimum_falling_path_sum

// dp数组不需要额外创建，在原数组中修改就可以了

// 除第一行外，数据变成上一行能到达该点的最小值加该点的值

// 最后返回最后一行的最小值即可

func minFallingPathSum(A [][]int) int {
	for i := 1; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			minN := A[i-1][j]
			if j > 0 {
				minN = min(minN, A[i-1][j-1])
			}
			if j < len(A[j])-1 {
				minN = min(minN, A[i-1][j+1])
			}
			A[i][j] = minN + A[i][j]
		}
	}
	res := A[len(A)-1][0]
	for j := 1; j < len(A); j++ {
		res = min(res, A[len(A)-1][j])
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
