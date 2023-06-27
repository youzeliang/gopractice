package _070_climbing_stairs

// https://leetcode-cn.com/problems/climbing-stairs/
//方法3 动态规划

// //dp原理：dp[i] = dp[i-1] + dp[i-2]
func climbStairs(n int) int {
	var tempMap = make(map[int]int)
	tempMap[1] = 1
	tempMap[2] = 2
	for i := 3; i <= n; i++ {
		tempMap[i] = tempMap[i-1] + tempMap[i-2]
	}

	return tempMap[n]
}

// 方法2 动态规划 + 优化
func climb4(n int) int {
	if n == 1 {
		return 1
	}
	one, two := 1, 2
	for i := 3; i <= n; i++ {
		one, two = two, one+two
	}

	return two
}
