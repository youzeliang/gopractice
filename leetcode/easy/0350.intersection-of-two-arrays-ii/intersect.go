package _350_intersection_of_two_arrays_ii

/**

https://leetcode.com/problems/intersection-of-two-arrays-ii/


给定两个数组，编写一个函数来计算它们的交集。


如果给定的数组已经排好序呢？你将如何优化你的算法？
如果 nums1 的大小比 nums2 小很多，哪种方法更优？
如果 nums2 的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
*/

func intersect(nums1 []int, nums2 []int) []int {
	res := []int{}
	m1 := getInts(nums1)
	m2 := getInts(nums2)

	if len(m1) > len(m2) {
		m1, m2 = m2, m1
	}
	// m1 是较短的字典，会快一些
	for n := range m1 {
		m1[n] = min(m1[n], m2[n])
	}

	for n, size := range m1 {
		for i := 0; i < size; i++ {
			res = append(res, n)
		}
	}

	return res

}

// 清理重复的值，也便于查询

func getInts(a []int) map[int]int {

	res := make(map[int]int, len(a))

	for i := range a {
		res[a[i]]++
	}
	return res
}

func min(a, b int) int {

	if a < b {
		return a
	}
	return b
}
