package _414_third_maximum_number

import (
	"math"
)

//
//# [414. Third Maximum Number](https://leetcode.com/problems/third-maximum-number/)
//
//## 题目
//
//Given a non-empty array of integers, return the third maximum number in this array. If it does not exist, return the maximum number. The time complexity must be in O(n).
//
//Example 1:
//```
//Input: [3, 2, 1]
//
//Output: 1
//
//Explanation: The third maximum is 1.
//```
//
//
//Example 2:
//```
//Input: [1, 2]
//
//Output: 2
//
//Explanation: The third maximum does not exist, so the maximum (2) is returned instead.
//```
//
//
//Example 3:
//```
//Input: [2, 2, 3, 1]
//
//Output: 1
//
//Explanation: Note that the third maximum here means the third maximum distinct number.
//Both numbers with value 2 are both considered as second maximum.
//```

func thirdMax(nums []int) int {
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64
	for _, n := range nums {
		if n == max1 || n == max2 { // 过滤掉前两大的重复数据
			continue
		}

		switch {
		case max1 < n:
			max3, max2, max1 = max2, max1, n
		case max2 < n:
			max3, max2 = max2, n
		case max3 < n:
			max3 = n
		}
	}

	// 说明没有第三大的数
	if max3 == math.MinInt64 {
		return max1
	}

	return max3
}

