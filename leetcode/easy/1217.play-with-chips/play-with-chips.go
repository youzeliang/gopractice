package _217_play_with_chips

// 移动偶数步无代价，移动奇数步代价一，找出数组中奇偶数的个数，将数量较少的一方移动奇数步（偶数-奇数=奇数 奇数-偶数=奇数）
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minCostToMoveChips(chips []int) int {
	oddCnt := 0
	evenCnt := 0
	for _, c := range chips {
		if c&1 != 0 {
			oddCnt++
		} else {
			evenCnt++
		}
	}
	return min(evenCnt, oddCnt)
}
