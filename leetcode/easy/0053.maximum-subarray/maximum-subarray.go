package _053_maximum_subarray

func maxSubArray(nums []int) int {
	sum, maxSum := -1<<31, -1<<31
	for _, n := range nums {
		// sum+n < n，那就还不如直接从 n 开始统计。
		sum = max(sum+n, n)
		maxSum = max(maxSum, sum)
	}
	return maxSum

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 解法2 动态规划

//dp =nums[0]   初始化
//dp =dp+nums[i]  dp>0
//dp = nums[i] dp<=0

func maxSubArray2(nums []int) int {
	curMaxSum := 0       // 当前最大连续子序列和
	allMaxSum := nums[0] // 全局最大连续子序列和
	for _, x := range nums {
		if curMaxSum < 0 { // 当前最大和如果小于0，直接用当前元素值覆盖这个负值
			curMaxSum = x
		} else { // 如果当前最大和>=0，继续累加
			curMaxSum += x
		}
		if curMaxSum > allMaxSum {
			allMaxSum = curMaxSum
		}
	}
	return allMaxSum
}

func maxSubArray3(nums []int) int {

	if len(nums) < 1 {
		return 0
	}

	dp := make([]int, len(nums))
	result := nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		result = max(dp[i], result)
	}
	return result
}
