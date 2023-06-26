package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	a := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}

	fmt.Println(len(a[0]))

}

type person struct {
	Name string
	Age  int
}

func ff() {
	p := person{Name: "我", Age: 20}
	//struct to json
	jsonB, err := json.Marshal(p)
	print(jsonB)
	if err == nil {
		fmt.Println(string(jsonB))
	}
	//json to struct
	respJSON := "{\"Name\":\"李四\",\"Age\":40}"
	json.Unmarshal([]byte(respJSON), &p)
	fmt.Println(p)
}

func nthUglyNumber(n int) int {

	if n == 0 {
		return 1
	}

	dp := make([]int, n)

	dp[0] = 1

	index2, index3, index5 := 0, 0, 0

	for i := 1; i < n; i++ {
		dp[i] = min(dp[index2]*2, dp[index3]*3, dp[index5]*5)

		if dp[i] == dp[index2]*2 {
			index2++
		}

		if dp[i] == dp[index3]*3 {
			index3++
		}

		if dp[i] == dp[index5]*5 {
			index5++
		}

	}

	return dp[n-1]
}

func min(a, b, c int) int {

	if a <= b && a <= c {
		return a
	}

	if b <= a && b <= c {
		return b
	}

	return c

}
