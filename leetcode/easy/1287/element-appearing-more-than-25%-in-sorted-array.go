package _287

// https://leetcode-cn.com/problems/element-appearing-more-than-25-in-sorted-array/
func findSpecialInteger(arr []int) int {
	l := len(arr)
	if l == 0 {
		return 0
	}

	res, count := 0, 0

	for i := 0; i < l; i++ {
		if count == 0 {
			res = arr[i]
			count++
		} else if arr[i] == res {
			count++
		} else {
			count--
		}

		if count > l/4 {
			break
		}

	}

	return res

}
