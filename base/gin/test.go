package main

import (
	"fmt"
	"sort"
)

func main() {

	c := findRepeatNumber([]int{1, 0, 2, 3, 5, 3})

	fmt.Println(c)
}

func canBeEqual(target []int, arr []int) bool {
	if len(target) != len(arr) {
		return false
	}
	sort.Sort(sort.IntSlice(target))
	sort.Sort(sort.IntSlice(arr))
	sort.Ints(target)

	for i := 0; i < len(target); i++ {
		if target[i] != arr[i] {
			return false
		}
	}

	return true
}

// 此题可以定义一个map来判断，但是有没有更好地方法不开辟新的空间来做，
// 由于数组中的元素都大于等于0，把对应位置的值变为负数
func findRepeatNumber(nums []int) int {
	for _, j := range nums {
		if j < 0 {
			j = -1 * j
		}
		if nums[j] < 0 {
			return j
		} else {
			nums[j] = -1 * j //设置为负数
		}

	}

	return 0 // 由于0没有办法变为负数，所以一旦走到这里，一定是0
}
