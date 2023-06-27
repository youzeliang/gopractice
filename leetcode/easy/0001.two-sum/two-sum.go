package problem0001

/*
*

https://leetcode.com/problems/two-sum/

给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

## 解题思路

```go
a + b = target
```

也可以看成是

```go
a = target - b
```

在`map[整数]整数的序号`中，可以查询到a的序号。这样就不用嵌套两个for循环了。
*/
func twoSum(nums []int, target int) []int {
	// index 负责保存map[整数]整数的序列号
	index := make(map[int]int, len(nums))

	for i, b := range nums {
		// 通过查询map，获取a = target - b的序列号
		if j, ok := index[target-b]; ok {
			return []int{j, i}
			// 注意，顺序是j，i
		}
		// 把b和i的值，存入map
		index[b] = i
	}

	return nil
}
