package _349_intersection_of_two_arrays

//# [349. Intersection of Two Arrays](https://leetcode.com/problems/intersection-of-two-arrays/)
//
//## 题目
//
//Given two arrays, write a function to compute their intersection.
//
//Example:
//Given nums1 = [1, 2, 2, 1], nums2 = [2, 2], return [2].
//
//Note:
//
//1. Each element in the result must be unique.
//1. The result can be in any order.

func intersection(nums1 []int, nums2 []int) []int {

	res := []int{}

	m1 := getInt(nums1)
	m2 := getInt(nums2)

	if len(nums1) > len(nums2) {
		m1, m2 = m2, m1
	}

	// m1 是较短的字典，会快一些
	for n := range m1 {
		if m2[n] {
			res = append(res, n)
		}
	}

	return res

}

func getInt(a []int) map[int]bool {

	res := make(map[int]bool, len(a))

	for i := range a {
		res[a[i]] = true
	}

	return res

}
