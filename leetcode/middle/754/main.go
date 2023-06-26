package _54

import "math"

func reachNumber(target int) int {
	t := math.Abs(float64(target))
	var sum = 0
	var forward = 0
	for sum < int(t) {
		forward++
		sum += forward
	}
	if (sum-int(t))%2 == 0 {
		return forward
	}
	return forward + 1 + (forward % 2)
}
