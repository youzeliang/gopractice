package _070_climbing_stairs

//方法3 动态规划
func climbStairs(n int) int {
	var tempMap = make(map[int]int)
	tempMap[1] = 1
	tempMap[2] = 2
	for i := 3; i <= n; i++ {
		tempMap[i] = tempMap[i-1] + tempMap[i-2]
	}

	return tempMap[n]
}
