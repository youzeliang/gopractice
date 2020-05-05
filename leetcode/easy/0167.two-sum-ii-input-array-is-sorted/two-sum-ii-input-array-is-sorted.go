package _167_two_sum_ii_input_array_is_sorted

func twoSum(numbers []int, target int) []int {
	// index 负责保存map[整数]整数的序列号
	index := make(map[int]int, len(numbers))

	for i, b := range numbers {
		// 通过查询map，获取a = target - b的序列号
		if j, ok := index[target-b]; ok {
			return []int{j + 1, i + 1}
			// 注意，顺序是j，i
		}
		// 把b和i的值，存入map
		index[b] = i
	}

	return nil
}
