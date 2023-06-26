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

<<<<<<< HEAD
		case nums[middle]> target:
			right = middle-1
=======
		case nums[middle] > target:
			right = middle - 1
<<<<<<< HEAD
>>>>>>> c8ee5da297c6341258cd96427f6b44501c49185a
=======
>>>>>>> 3955728f17484cc4ab190a6b46fde0df977ec4e4
>>>>>>> 35c5e937bc81bbbc8299e6c42daff1e83721e0f3
		default:
			return middle
		}
	}

	m := make(map[int]struct{})

<<<<<<< HEAD
	return - 1
=======
	return -1
<<<<<<< HEAD
>>>>>>> c8ee5da297c6341258cd96427f6b44501c49185a
=======
>>>>>>> 3955728f17484cc4ab190a6b46fde0df977ec4e4
>>>>>>> 35c5e937bc81bbbc8299e6c42daff1e83721e0f3
}
