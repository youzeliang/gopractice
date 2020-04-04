package _198_house_robber

// https://leetcode-cn.com/problems/house-robber/
func rob(nums []int) int {
	// a 对于偶数位上的最大值的记录
	// b 对于奇数位上的最大值的记录
	var a, b int
	for i, v := range nums {
		if i%2 == 0 {
			a = max(a+v, b)
		} else {
			b = max(a, b+v)
		}
	}

	return max(a, b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//假设现在有四个元素，[6,7,3,4]，前三个元素没得选，6、7不变，3只能加6。轮到4时，4应该 4+7 还是 4+6？
//因为数组是无序的，对每个数字，显然应该比较他能取到的两个数哪个更大。
//用dp[i]来保存第i位时的最大值。
// dp[i] = max(dp[i-3]+nums[i] , dp[i-2]+nums[i])
func rob1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	if len(nums) == 1 {
		return dp[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	if len(nums) == 3 {
		return max(nums[0]+nums[2], nums[1])
	}

	dp[1] = nums[1]
	dp[2] = nums[2] + nums[0]
	for i := 3; i < len(nums); i++ {
		dp[i] = max(dp[i-3]+nums[i], dp[i-2]+nums[i])
	}

	return max(dp[len(nums)-1], dp[len(nums)-2])

}

// 使用临时变量来解决
func rob2(nums []int) int {
	preMax, curMax := 0, 0
	for _, m := range nums {
		curMax, preMax = max(preMax+m, curMax), curMax
	}
	return curMax
}
