package _035_search_insert_position

//# [35. Search Insert Position](https://leetcode.com/problems/search-insert-position/)
//
//## 题目
//Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
//
//You may assume no duplicates in the array.
//```
//Here are few examples.
//[1,3,5,6], 5 → 2
//[1,3,5,6], 2 → 1
//[1,3,5,6], 7 → 4
//[1,3,5,6], 0 → 0
//```

func searchInsert(nums []int, target int) int {
	// 没有把i放入for语句中
	// 是为了兼容，len(nums) == 0 和 target > nums[len(nums)-1]两种情况
	i := 0

	for i < len(nums) && nums[i] <= target {
		// 相等的时候，直接返回
		if nums[i] == target {
			return i
		}

		// 否则，就去检查下一个
		i++
	}

	return i
}
