package problem0217

// [217. Contains Duplicate](https://leetcode.com/problems/contains-duplicate/)

/*

给定一个整数数组，判断是否存在重复元素。

如果任何值在数组中出现至少两次，函数返回 true。如果数组中每个元素都不相同，则返回 false。

*/

func containsDuplicate(nums []int) bool {
	// 利用 m 记录 nums 中的元素是否出现过
	m := make(map[int]bool, len(nums))
	for _, n := range nums {
		if m[n] {
			return true
		}
		m[n] = true
	}

	return false

}
