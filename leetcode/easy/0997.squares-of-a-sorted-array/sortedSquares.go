package _997_squares_of_a_sorted_array

import "sort"

// https://leetcode-cn.com/problems/squares-of-a-sorted-array/

func sortedSquares(A []int) []int {
	for i := 0; i < len(A); i++ {
		A[i] = A[i] * A[i]
	}

	sort.Ints(A)
	return A
}
