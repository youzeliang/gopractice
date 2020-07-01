package _523_continuous_subarray_sum

func checkSubarraySum(nums []int, k int) bool {
	l := len(nums)
	hashmap := make(map[int]int)
	hashmap[0] = -1 //对0特殊处理
	sum := 0
	for i := 0; i < l; i++ {
		sum += nums[i]
		if k != 0 {
			sum = sum % k
		}

		if v, ok := hashmap[sum]; ok {
			if i-v > 1 { //必须有2个数
				return true
			}
		} else {
			hashmap[sum] = i
		}
	}
	return false

}
