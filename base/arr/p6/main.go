package main

func main() {

	//J := "aA"
	//S := "aAAbbbb"
	//
	//numJewelsInStones(J,S)

	findMaxLength([]int{0, 1, 0, 1})

}

func findMaxLength(nums []int) int {
	indices := map[int]int{0: -1}
	var res, prefixSum int
	for i, v := range nums {
		if v == 0 {
			v = -1
		}
		prefixSum += v
		if j, ok := indices[prefixSum]; !ok {
			indices[prefixSum] = i
		} else if i-j > res {
			res = i - j
		}
	}
	return res
}

func numJewelsInStones(J string, S string) int {
	isJewel := make(map[byte]bool, len(J))
	for i := range J {
		isJewel[J[i]] = true
	}

	res := 0
	for i := range S {
		if isJewel[S[i]] {
			res++
		}
	}

	return res
}
