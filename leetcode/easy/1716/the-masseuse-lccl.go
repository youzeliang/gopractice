package _716

// https://leetcode-cn.com/problems/the-masseuse-lcci/

// 解题思路
// 状态定义：状态定义：首先考虑只有nums[0],nums[1],nums[2]三个下标时，又因为时间都是正数，所以最长时间为max(nums[1],nums[0]+nums[2])，dp[n-1]就是前一天的最长预约时间与前二天的最长预约时+当天预约时间较大者
// dp公式：dp[i] = max(dp[i-1],dp[i-2] + nums[i])

func massage(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	if l == 1 {
		return nums[0]
	}

	dp := make([]int,l+1)
	dp[0] = nums[0]
	dp[1] = max(nums[0],nums[1])

	for i := 2;i < l ; i++ {
		dp[i] = max(dp[i-1],dp[i-2]+nums[i])

	}

	return dp[l-1]
}

func max(a, b int ) int {
	if a > b {
		return a
	}
	return b
}
