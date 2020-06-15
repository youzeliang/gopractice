package _413_arithmetic_slices

// N = n的手，数量为：n+(n-1)+(n-2)+...+3+2+1 个
func numberOfArithmeticSlices(A []int) int {
	l := len(A)
	if l <= 2 {
		return 0
	}
	r := 0
	curSliceNum := 0 // 当前连续数
	for i := 2; i < l; i++ {
		if A[i-1]-A[i-2] == A[i]-A[i-1] {
			curSliceNum++
		} else {
			curSliceNum = 0
		}
		r += curSliceNum
	}
	return r
}
