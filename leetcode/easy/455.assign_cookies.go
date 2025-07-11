package easy

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	i, j := 0, 0

	for i < len(g) && j < len(s) {
		if g[i] <= s[i] {
			i++
		}
		j++
	}
	return i
}
