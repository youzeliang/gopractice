package _026_remove_duplicates_from_sorted_array

/**
https://leetcode.com/problems/remove-duplicates-from-sorted-array/


题目

给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。


给定数组 nums = [1,1,2],

函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。

你不需要考虑数组中超出新长度后面的元素。



解题思路

遍历数组，只有后面的数比前面的数大，就说明前面的元素出现了一次。然后做一个累加值

因为是排序好了的数组，所以 num[n+1] >= num[n]


*/

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i, j := 1, 0
	for ; i < len(nums); i++ {
		if nums[i] > nums[j] {
			j++
			nums[j] = nums[i]
		}
	}
	j++

	return j
}

func removeDuplicates2(nums []int) int {
	for i := 0; i+1 < len(nums); {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}

	return len(nums)
}
