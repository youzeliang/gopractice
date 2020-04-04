package _657_judge_route_circle

import "strings"

func judgeCircle(moves string) bool {
	x, y := 0, 0
	for i := 0; i < len(moves); i++ {
		switch moves[i] {
		case 'R':
			x++
		case 'L':
			x--
		case 'U':
			y++
		case 'D':
			y--
		}

	}

	if x == 0 && y == 0 {
		return true
	}

	return false
}

// 解法2，利用内置函数

func judgeCircle2(moves string) bool {
	up := strings.Count(moves, "U")
	down := strings.Count(moves, "D")
	left := strings.Count(moves, "L")
	right := strings.Count(moves, "R")

	return up == down && left == right
}
