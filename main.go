package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := []int{2, 2, 3, 1}

	sort.Ints(nums)

	for i, j := range nums {

	}

	if len(nums) <= 2 {
		fmt.Println(nums[len(nums)-3])
	}

	fmt.Println(nums[len(nums)-3])

}
