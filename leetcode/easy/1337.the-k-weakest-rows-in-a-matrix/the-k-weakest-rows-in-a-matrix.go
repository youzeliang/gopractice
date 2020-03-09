package _337_the_k_weakest_rows_in_a_matrix

import "sort"

// https://leetcode-cn.com/problems/the-k-weakest-rows-in-a-matrix/
// 统计每行多少个 1
// 按每行的 1 个数从小到大排序
func kWeakestRows(mat [][]int, k int) []int {
	m, n := len(mat), len(mat[0])
	res := make([]int, m)
	ans := make([]int, m)
	for i := 0; i < m; i++ {
		ans[i] = i
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				res[i]++
			}
		}
	}

	sort.Slice(ans, func(i, j int) bool {
		return res[ans[i]] < res[ans[j]] || (res[ans[i]] == res[ans[j]] && ans[i] < ans[j])
	})

	return ans[:k]
}
