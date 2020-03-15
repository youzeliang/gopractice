package _7

import (
	"math"
)

func printNumbers(n int) []int {
	var res []int
	l := int(math.Pow(10, float64(n))) - 1
	for i := 1; i <= l; i++ {
		res = append(res, i)
	}
	return res

}
