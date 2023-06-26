package main

import "fmt"

func main() {

	var c = []int{1, 2, 3, 4, 5}
	i := 2
	c = append(c[:i], c[i+1:]...)
	fmt.Println(c)
	findDisappearedNumbers(c)

}

// 遍历数组，将每个数字交换到它理应出现的位置上，下面情况不用换：
// 当前数字本就出现在理应的位置上，跳过，不用换。
// 当前数字理应出现的位置上，已经存在当前数字，跳过，不用换。
// 再次遍历，如果当前位置没对应正确的数，如上图索引 4、5，则将对应的 5、6 加入 res。

func findDisappearedNumbers(nums []int) []int {
	i := 0
	for i < len(nums) {
		if nums[i] == i+1 { // 当前元素出现在它该出现的位置，无需交换
			i++
			continue
		}
		idealIdx := nums[i] - 1        // idealIdx：当前元素应该出现的位置
		if nums[i] == nums[idealIdx] { // 当前元素=它理应出现的位置上的现有元素，说明重复了
			i++
			continue
		}
		nums[idealIdx], nums[i] = nums[i], nums[idealIdx] // 不重复，进行交换
		// 这里不要i++，因为交换过来的数字本身也需要考察，需要交换到合适的位置上
		// 如果 i++ 就会跳过它，少考察了它
	}

	res := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 { // 值与索引 不对应
			res = append(res, i+1)
		}
	}
	return res
}
