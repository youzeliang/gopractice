package _004_median_of_two_sorted_arrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	return 0
}

// 合并两个排序好的数组
//func combine(mis, njs []int) []int {
//	lenM, m := len(mis), 0
//	lenN, n := len(njs), 0
//	res := make([]int, lenN+lenM)
//
//	for k := 0; k < lenM+lenN; k++ {
//
//	}
//
//}

// 求出中位数
func median(nums []int) float64 {
	l := len(nums)

	if l == 0 {
		panic("切片的长度为0，无法求解中位数。")
	}

	if l%2 == 0 {
		return float64(nums[l/2]+nums[l/2+1]) / 2.0
	}

	return float64(nums[l/2])
}
