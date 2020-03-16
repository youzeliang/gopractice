package _365_how_many_numbers_are_smaller_than_the_current_number

// https://leetcode-cn.com/problems/how-many-numbers-are-smaller-than-the-current-number/

// 暴力法
func smallerNumbersThanCurrent(nums []int) []int {

	res := make([]int, 0)

	for i, j := range nums {
		r := 0
		for m, n := range nums {
			if m != i && j > n {
				r++
			}
		}
		res = append(res, r)
	}

	return res
}

// O(S + n) 大佬的方法
func smallerNumbersThanCurrent2(nums []int) []int {
	count := make([]int, 101)
	res := make([]int, len(nums))

	for _, v := range nums {
		count[v]++
	}

	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	for i, v := range nums {
		if v != 0 {
			res[i] = count[v-1]
		}
	}

	return res
}
