package main

import "strconv"

func main() {
	monotoneIncreasingDigits(332)
}

func monotoneIncreasingDigits(N int) int {
	if N < 10 {
		return N
	}
	res := []byte(strconv.Itoa(N))
	length := len(res)
	for i := 0; i < length-1; {
		//当前数比后一位大，当前就减一，可能造成前一位比当前位大，所以i 后退
		if res[i] > res[i+1] {
			res[i]--
			for j := i + 1; j < length; j++ {
				res[j] = '9'
			}
			i--        //i后退
			if i < 0 { //第一位就跳出去
				break
			}
		} else {
			i++
		}
	}
	sum, _ := strconv.Atoi(string(res))
	return sum
}
