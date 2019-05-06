package problem0001

func twoSum(nums []int, target int) []int {
	// index 负责保存map[整数]整数的序列号
	index := make(map[int]int, len(nums))

	for i, b := range nums {
		// 通过查询map，获取a = target - b的序列号
		if j, ok := index[target-b]; ok {
			return []int{j, i}
			// 注意，顺序是j，i
		}

		index[b] = i
	}

	return nil
}
