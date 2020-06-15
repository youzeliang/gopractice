package _416_partition_equal_subset_sum

//将此题题意转换为：背包容量为sum/2 物品为题意给出的数组 问能否取出一次物品恰好装满背包 将此问题转化为0-1背包问题
//状态：背包容量以及当前物品
//选择：若当前背包容量大于当前要去的物品重量则可将其装进或不装进即可
//若当前背包容量小于当前要取的物品的重量则当前dp状态为之前的状态
//转化方程：dp[j]=dp[j]||dp[j-nums[i]]

func canPartition(nums []int) bool {

	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%2 == 1 {
		return false
	}

	sum /= 2
	dp := make([]bool, sum+1)
	dp[0] = true

	for i := 0; i < len(nums); i++ {
		for j := sum; j >= 0; j-- {
			if j-nums[i] >= 0 {
				dp[j] = dp[j] || dp[j-nums[i]]
			}

		}
	}
	return dp[sum]

}
