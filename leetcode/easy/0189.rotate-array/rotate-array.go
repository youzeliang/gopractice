package _189_rotate_array

/**

https://leetcode.com/problems/rotate-array/


给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。


输入: [1,2,3,4,5,6,7] 和 k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]



解题思路

如果 n = 7 , k =3
把原数组划分为两部分来看，前n-k 个元素[1,2,3,4] 和后k 个元素 [5,6,7]

1.定义 reverse 逆转方法：将数组元素反转，比如 [1,2,3,4] 逆转后变成 [4,3,2,1]
2.对前 n - k 个元素 [1,2,3,4] 进行逆转后得到 [4,3,2,1]
3. 对后k个元素 [5,6,7] 进行逆转后得到 [7,6,5]
4. 将前后元素 [4,3,2,1,7,6,5] 逆转得到：[5,6,7,1,2,3,4]


注意：还要处理 k > 数组长度的情况，对 k 进行取模


*/

func rotate(nums []int, k int) {

	n := len(nums)

	k %= n

	if k == 0 || k == n {
		return
	}

	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)

}

func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
