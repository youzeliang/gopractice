package _096_unique_binary_search_trees

//卡特兰数 公式
//G(n) = G(0)*G(n-1)+G(1)*G(n-2)+...+G(n-1)*G(0)

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

func numTrees1(n int) int {
	var res = 1
	for i := 0; i < n; i++ {
		res = res * 2 * (1 + 2*i) / (i + 2)
	}
	return res
}
