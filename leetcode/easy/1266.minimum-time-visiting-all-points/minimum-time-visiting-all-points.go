package _266_minimum_time_visiting_all_points

import "math"

func minTimeToVisitAllPoints(points [][]int) int {
	var sum int = 0
	for i := 0; i < len(points)-1; i++ {
		x := int(math.Abs(float64(points[i+1][0] - points[i][0])))
		y := int(math.Abs(float64(points[i+1][1] - points[i][1])))
		temp := x
		if y > x {
			temp = y
		}
		sum += temp
	}
	return sum
}
