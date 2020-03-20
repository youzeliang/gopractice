package _315_count_of_smaller_numbers_after_self

import (
	"math"
	"sort"
)

func countSmaller(nums []int) []int {
	n := len(nums)

	var res []int
	for i := 0; i < n; i++ {
		tem := 0
		for j := i; j < n; j++ {
			if nums[j] < nums[i] {
				tem++

			}
		}
		res = append(res, tem)
	}

	return res
}

func luckyNumbers(matrix [][]int) []int {
	var res []int

	l := len(matrix)
	for i := 0; i < l; i++ {
		min := math.MaxInt32
		m := math.MaxInt32
		tem := 0
		for j := 0; i < l; j++ {

			if matrix[i][j] < m {
				tem = matrix[i][j]
				m = matrix[i][j]

			}

			if tem < min {
				min = tem
			}
		}
		res = append(res, min)

	}

	sort.Ints(res)

	return []int{res[len(res)-1]}
}
