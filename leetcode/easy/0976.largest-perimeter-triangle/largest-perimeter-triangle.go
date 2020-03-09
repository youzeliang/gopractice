package _976_largest_perimeter_triangle

import "sort"

//# [976. Largest Perimeter Triangle](https://leetcode.com/problems/largest-perimeter-triangle/)
//
//Given an array A of positive lengths, return the largest perimeter of a triangle with non-zero area, formed from 3 of these lengths.
//
//If it is impossible to form any triangle of non-zero area, return 0.
//
//Example 1:
//
//```text
//Input: [2,1,2]
//Output: 5
//```
//
//Example 2:
//
//```text
//Input: [1,2,1]
//Output: 0
//```
//
//Example 3:
//
//```text
//Input: [3,2,3,4]
//Output: 10
//```
//
//Example 4:
//
//```text
//Input: [3,6,2,3]
//Output: 8
//```
//
//Note:
//
//- 3 <= A.length <= 10000
//- 1 <= A[i] <= 10^6

func largestPerimeter(A []int) int {
	sort.Ints(A)
	for i := len(A) - 1; i-2 >= 0; i-- {
		if A[i] < A[i-1]+A[i-2] {
			return A[i] + A[i-1] + A[i-2]
		}
	}
	return 0
}
