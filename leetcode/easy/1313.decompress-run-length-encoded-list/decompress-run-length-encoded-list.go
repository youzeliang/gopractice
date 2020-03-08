package _313_decompress_run_length_encoded_list

// https://leetcode-cn.com/problems/decompress-run-length-encoded-list/
func decompressRLElist(nums []int) []int {
	var res []int
	for i := 0; i < len(nums); i += 2 {
		for j := 0; j < nums[i]; j++ {
			res = append(res, nums[i+1])
		}
	}

	return res
}
