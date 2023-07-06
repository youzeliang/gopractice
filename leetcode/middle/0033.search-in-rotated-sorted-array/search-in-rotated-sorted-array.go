package _033_search_in_rotated_sorted_array

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		// 判断是否在前半部分查找
		if (nums[left] <= target && target <= nums[mid]) || (nums[mid] <= nums[right] && (target < nums[mid] || target > nums[right])) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

// 参考 https://leetcode-cn.com/problems/search-in-rotated-sorted-array/solution/jian-dan-luo-lie-mei-yi-chong-qing-kuang-by-liu-zh/

func sea(nums []int, target int) int {

	left := 0
	right := len(nums) - 1

	for left < right {

		middle := (left + right) / 2
		if nums[middle] == target {
			return middle
		}

		if (nums[left] <= target && target <= nums[middle]) || (nums[middle] <= nums[right] && target <= nums[middle]) {
			right = middle - 1
		} else {
			left = middle + 1
		}

	}

	return -1

}

func s(nums []int, target int) int {

	left, right := 0, len(nums)-1

	for left < right {
		middle := (left + right) / 2

		switch {
		case nums[middle] < target:
			left = middle + 1

		case nums[middle] > target:
			right = middle - 1
		case nums[middle] > target:
			right = middle - 1
		default:
			return middle
		}
	}

	return -1
}
