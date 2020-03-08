package _922_sort_array_by_parity_ii

//
//# [922. Sort Array By Parity II](https://leetcode.com/problems/sort-array-by-parity-ii/)
//
//Given an array `A` of non-negative integers, half of the integers in A are odd, and half of the integers are even.
//
//Sort the array so that whenever `A[i]` is odd, `i` is odd; and whenever `A[i]` is even, `i` is even.
//
//You may return any answer array that satisfies this condition.
//
//Example 1:
//
//```text
//Input: [4,2,5,7]
//Output: [4,5,2,7]
//Explanation: [4,7,2,5], [2,5,4,7], [2,7,4,5] would also have been accepted.
//```
//
//Note:
//
//1. `2 <= A.length <= 20000`
//1. `A.length % 2 == 0`
//1. `0 <= A[i] <= 1000`

func sortArrayByParityII(A []int) []int {
	size := len(A)
	res := make([]int, size)
	even, odd := 0, 1

	for _, a := range A {
		if a%2 == 0 {
			res[even] = a
			even += 2
		} else {
			res[odd] = a
			odd += 2
		}
	}

	return res
}

// 不使用额外空间的情况下
func sortArrayByParityII2(A []int) []int {
	j := 1
	for i := 0; i < len(A); i += 2 {
		if A[i]%2 == 1 { //j
			for A[j]%2 == 1 {
				j += 2
			}
			//j 是o
			t := A[i]
			A[i] = A[j]
			A[j] = t
		}

	}
	return A

}
