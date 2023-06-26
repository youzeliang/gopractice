package _448_find_all_numbers_disappeared_in_an_array

import (
	"fmt"
	"testing"
)

func Test_findDisappearedNumbers(t *testing.T) {

	var a = []int{1, 4, 1, 2}
	find(a)
}

func find(nums []int) {
	for _, value := range nums {
		if value < 0 {
			value = -value
		}
		if nums[value-1] > 0 {
			nums[value-1] = -nums[value-1]
		}
	}

	fmt.Println(nums)
}
