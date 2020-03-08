package _801

// 第n层楼梯的状态转移方程为 dp[n] = (dp[n-3] + dp[n-2] + dp[n-1])%(1e9+7)
// https://leetcode-cn.com/problems/three-steps-problem-lcci/

func waysToStep(n int) int {
	dp := make([]int, n+1)
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 4
	}

	dp[0], dp[1], dp[2], dp[3] = 0, 1, 2, 4
	for i := 4; i < n+1; i++ {
		dp[i] = (dp[i-1] + dp[i-2] + dp[i-3]) % 1000000007
	}

	return dp[n]
}
