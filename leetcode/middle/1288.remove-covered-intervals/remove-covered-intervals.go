package _288_remove_covered_intervals

func removeCoveredIntervals(intervals [][]int) int {
	n := len(intervals)
	res := n
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && intervals[j][0] <= intervals[i][0] && intervals[i][1] <= intervals[j][1] {
				res--
				break
			}
		}
	}

	return res

}
