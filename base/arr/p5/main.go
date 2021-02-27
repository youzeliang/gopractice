package main

func main() {

	var a = []int{8,1,2,2,3}
	smallerNumbersThanCurrent(a)

}

func smallerNumbersThanCurrent(nums []int) []int {
	count := make([]int, 101)
	res := make([]int, len(nums))

	//
	for _, v := range nums {
		count[v]++
	}

	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	for i, v := range nums {
		if v != 0 {
			res[i] = count[v-1]
		}
	}

	return res
}
