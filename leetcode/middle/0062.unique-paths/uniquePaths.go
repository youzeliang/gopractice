package _062_unique_paths

func uniquePaths(m int, n int) int {
	// path[i][j] 代表了，到达 (i,j) 格子的不同路径数目
	path := [100][100]int{}

	for i := 0; i < m; i++ {
		// 到达第 0 列的格子，只有一条路径
		path[i][0] = 1
	}

	for j := 0; j < n; j++ {
		// 到达第 0 行的格子，只有一条路径
		path[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 到达 (i,j) 格子的路径数目，等于
			// 到达 上方格子 和 左边格子 路径数之和
			path[i][j] = path[i-1][j] + path[i][j-1]
		}
	}

	return path[m-1][n-1]
}

/**
1. 定义状态 dp[i][j] 为当前 i*j 的路径总数
2. 递推方程 dp[i][j] = dp[i-1][j]+dp[i][j-1]
3. 状态初始化 dp[i][1]=1 dp[1][j]=1
4. 最终结果  dp[m-1][n-1]
5. 优化每次仅用到 dp[i-1][j] 和 dp[i][j-1] 所以可将dp优化为一个一维数组中:cur[j] += cur[j-1]
    a) cur 长度取 n 防止溢出
    b) 方程为 cur[j] += cur[j-1] 等价于 cur[j] = cur[j] + cur[j-1] cur[i]每行遍历只更新一次,也就是说cur[j]原有的值代表的是上一行p[i-1],[j]的值,而cur[j-1] 是在本行遍历刚更新过的值因此是dp[i][j-1]的值!
*/

func uniquePaths1(m int, n int) int {
	cur := make([]int, n)
	for i := 0; i < n; i++ {
		cur[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			cur[j] += cur[j-1]
		}
	}
	return cur[n-1]
}
