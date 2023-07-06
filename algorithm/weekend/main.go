package main

import (
	"fmt"
	//"github.com/spf13/cast"
	"sort"
	"strings"
)

const (
	FromTypeDixData = 7
)

const (
	Services uint8 = iota
	PC
	Ipad
	IPhone
	Android
	PCClient
	Touch
	FixData
)

func main() {

	var a = []int{6, 6, 6, 6, 0, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6}

	s := 0
	for i := 0; i < len(a); i++ {
		if a[i] == 6 {
			s++
		}
	}

	fmt.Println(s, "ssssss")

	var buffer strings.Builder

	aa := []int{1, 2, 3, 4}
	for i := 0; i < len(aa); i++ {
		buffer.WriteString(string(i))
		buffer.WriteString(",")
	}

	fmt.Println(buffer)

	//var a = "123"

	//minimumRounds(a)

}

func minimumRounds(tasks string) {

	fmt.Println(string(FixData))

}
func tupleSameProduct(nums []int) int {
	if len(nums) > 4 {
		nums = append(nums, nums[:len(nums)-1]...)
	}

	res := 0
	for i := 0; i <= len(nums)-4; i++ {
		if c(nums[i : i+4]) {
			res++
		}
	}

	return res * 8

}

func c(nums []int) bool {
	sort.Ints(nums)
	return nums[0]*nums[3] == nums[1]*nums[2]
}

// 输入：nums = [2,3,4,6,8,12]
