package _977_squares_of_a_sorted_array

// 解法1
//func sortedSquares(A []int) []int {
//	arr := make([]int, 0)
//	for _, v := range A {
//		v *= v
//		arr = append(arr, v)
//	}
//	sort.Ints(arr)
//
//    return arr
//}

// 解法2
func sortedSquares(A []int) []int {
	size := len(A)
	res := make([]int, size)
	for l, r, i := 0, size-1, size-1; l <= r; i-- {
		if A[l]+A[r] < 0 {
			res[i] = A[l] * A[l]
			l++
		} else {
			res[i] = A[r] * A[r]
			r--
		}
	}
	return res
}
