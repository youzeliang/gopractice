package _905_sort_array_by_parity

import "math"

//# [905. Sort Array By Parity](https://leetcode.com/problems/sort-array-by-parity/)
//
//## 题目
//
//Given an array `A` of non-negative integers, return an array consisting of all the even elements of `A`, followed by all the odd elements of `A`.
//
//You may return any answer array that satisfies this condition.
//
//Example 1:
//
//```text
//Input: [3,1,2,4]
//Output: [2,4,3,1]
//The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.
//```
//
//Note:
//
//- 1 <= A.length <= 5000
//- 0 <= A[i] <= 5000

func sortArrayByParity(A []int) []int {
	newArray := make([]int, 0)

	for _, j := range A {
		if j%2 == 0 {
			newArray = append(newArray, j)
		}
	}

	for _, m := range A {
		if m%2 != 0 {
			newArray = append(newArray, m)
		}
	}
	var a = math.MaxInt32

	return newArray
}

// 只要把奇数和偶数分开就行。假定把偶数放在前面，初始一个起始位置”c”，遍历数组“A”如果当前值为偶数即和“c”的值交换，并且“c"本身加一。
func sortArrayByParity2(A []int) []int {
	c := 0
	for i := range A {
		if A[i]%2 == 0 {
			A[c], A[i] = A[i], A[c]
			c++
		}
	}
	return A
}
