package _232_check_If_It_Is_a_straight_line

import "sort"

//https: //leetcode-cn.com/problems/check-if-it-is-a-straight-line/solution/golang-by-resara-13/
// 用2点来确定一条直线，然后判断数组的其他元素是否在这条直线上,不需要求出斜率，因为可能是负数，用交叉相乘来判断

func checkStraightLine(coordinates [][]int) bool {

	c := coordinates
	x, y := c[1][0]-c[0][0], c[1][1]-c[0][1]

	for i := 2; i < len(c); i++ {
		x1, y1 := c[i][0]-c[i-1][0], c[i][1]-c[i-1][1]
		if x1*y != x*y1 {
			return false
		}
	}

	a := make([]int, 1)

	sort.Ints(a)
	return true

}
