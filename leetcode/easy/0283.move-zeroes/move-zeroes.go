package _283_move_zeroes

// 先遍历数组，将不为0的值填充到数组中，再遍历一次进行补0

func moveZeroes(nums []int) {
	index := 0
	for _, v := range nums {
		if v != 0 {
			nums[index] = v
			index++
		}
	}
	for i := index; i < len(nums); i++ {
		nums[i] = 0
	}

}
