package main

import (
	"bytes"
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

type User struct {
	UserId   int
	UserName string
}
type UserTag struct {
	UserId   int    `json:"user_id" bson:"user_id"`
	UserName string `json:"user_name" bson:"user_name"`
}

const (
	a = iota
	b = iota
)
const (
	name     = "name"
	fdfsdfsd = "fafa"
	c        = iota
	d        = iota
)

func generateTheString(n int) string {

	a := make([]string, 0)
	if n%2 == 0 {
		for i := 0; i < n-1; i++ {
			a = append(a, "a")
		}
		a = append(a, "b")

	} else {
		for i := 0; i < n-1; i++ {
			a = append(a, "a")
		}

	}

	return strings.Join(a, "")

}

func numTimesAllBlue(light []int) int {
	sumA := 0
	sumB := 0
	res := 0
	for i, j := range light {
		sumA = sumA + 1 + i
		sumB += j
		if sumA == sumB {
			res++
		}
	}

	return res
}

func smallerNumbersThanCurrent2(nums []int) {
	count := make([]int, 6)
	//res := make([]int, len(nums))

	for _, v := range nums {
		count[v]++
	}

	println("-----------")
	for _, j := range count {
		println(j)
	}

	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	println("========")
	for i := 0; i < len(count); i++ {
		println(count[i])
	}

	println("1111", count)

	//for i, v := range nums {
	//	if v != 0 {
	//		res[i] = count[v-1]
	//	}
	//}
	//
	//return res
}

func main() {

	//smallerNumbersThanCurrent2([]int{3, 4, 5})
	//majorityElement([]int{1, 1, 2, 3})
	//ta()

	//sumFourDivisors([]int{7, 9})
	//fdsfds()

	//findLucky([]int{2, 2, 3, 4})

}

func ginInfo() {

	gin.Default()
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run() // 监听并在 0.0.0.0:8080 上启动服务

}

func findLucky(arr []int) int {
	m := make(map[int]int)
	for _, j := range arr {
		if _, ok := m[j]; ok {
			m[j] = m[j] + 1
		} else {
			m[j] = 1
		}
	}

	if len(m) == 0 {
		return -1
	}

	max := 0
	for i, j := range m {
		if i == j {
			if j > max {
				max = j
			}
		}
	}

	return max
}

func fdsfdddds() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[0:3] // 左闭右开区间，最终切片为 [1,2,3]
	println(cap(slice))

	m := make(map[int]int)
	if _, ok := m[1]; ok {

	}
}

func sumFourDivisors(nums []int) int {
	var res int
	for i := 0; i < len(nums); i++ {
		res += lll(nums[i])
	}

	return res

}

func lll(num int) int {

	if num < 6 {
		return 0
	}

	res := 0
	sum := 0
	n := int(math.Sqrt(float64(num)))

	re := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			res++
			sum += i
			re = i
		}
		if res > 2 {
			return 0
		}
	}

	if res == 2 {
		return 1 + num + re + num/re
	}

	return 0

}

func ta() {
	n := 1
	println(int(math.Pow(10, float64(n))))
}

func printNumbers(n int) []int {
	var res []int
	l := int(math.Pow(10, float64(n))) - 1
	for i := 1; i <= l; i++ {
		res = append(res, i)
	}
	return res
}

func ttt() {
	//res := make([]int, 5)

	var res = []int{1, 1, 2, 2, 3, 4, 5}

	for i := 0; i < 7; i++ {
		res[i]++
	}

	for _, j := range res {
		println("-----", j)
	}
}

func majorityElement(nums []int) int {
	numLens := len(nums)
	res := make(map[int]int, numLens)
	for _, v := range nums {
		res[v]++
	}

	for _, j := range res {
		println("-----", j)
	}

	return -1
}

func arrayPairSum(nums []int) int {
	res := 0

	for i := 1; i < len(nums)+1; i += 2 {
		if nums[i] > nums[i-1] {
			res += nums[i-1]
		} else {
			res += nums[i]
		}
	}

	return res
}

func rankTeams(votes []string) string {
	if len(votes) == 0 || len(votes) == 1 {
		return strings.Join(votes, ",")
	}

	a := make(map[int]int, 0)

	// 每一个字母得多少分
	for i := range votes {
		a[i] = len(votes) - i
	}

	v := make(map[int]int, 0)
	for _, j := range votes {
		for m := range j {
			if _, ok := a[int(j[m])]; ok {
				if _, ok := v[int(j[m])]; ok {
					v[int(j[m])] = int(j[m]) + a[int(j[m])]
				} else {
					v[int(j[m])] = a[int(j[m])]
				}
			}
		}
	}

	x := make(map[int]int, 0)
	y := make(map[int]int, 0)
	for _, j := range v {
		if _, ok := x[j]; ok {
			y[j] = j
		}
		x[j] = j
	}

	g := make([]int, 0)
	for _, n := range v {
		g = append(g, n)
	}

	sort.Ints(g)

	for i, n := range v {
		if _, ok := v[g[i]]; ok {

		}

		fmt.Print(n)

		if i >= 1 && v[i-1] == v[i] {

		}
	}
	return ""

}

func tes11() {
	s := []string{"ewq", "ds"}

	for _, j := range s {
		//println(s[j])

		println(j)
		//for m := range j {
		//	//println(j[m])
		//}

	}
}

func closestDivisors(num int) []int {
	a := math.Sqrt(float64(num + 1))
	b := math.Sqrt(float64(num + 2))

	var res []int
	a1 := int(a)

	min := float64(math.MaxInt32)
	resa := 0
	resb := 0
	if num == 1 {
		return append(res, 1, 2)
	}

	if num == 2 {
		return append(res, 2, 2)
	}

	for i := 1; i <= num; i++ {
		if (num + 1) == i*a1 {
			if math.Abs(float64(i-a1)) < min {
				resa = i
				resb = a1
				min = math.Abs(float64(i - a1))
			}
			a1--
		}
	}

	b1 := int(b)

	for i := 1; i <= num; i++ {
		if (num + 2) == i*b1 {
			if math.Abs(float64(i-b1)) < min {
				resa = i
				resb = b1
				min = math.Abs(float64(i - b1))
			}
			b1--
		}
	}

	return append(res, resb, resa)

}

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	var res byte = 0

	for i := range s1 {
		res ^= s1[i]
	}

	for i := range s2 {
		res ^= s2[i]
	}

	return res == 0
}

func minSteps(s string, t string) int {

	i := 0
	for _, j := range t {
		if !strings.Contains(string(j), s) {
			i++
		}
	}

	return len(s) - i
}

func isIsomorphic(s string, t string) bool {
	sb := []byte(s)
	st := []byte(t)

	if len(sb) == 0 && len(st) == 0 {
		return true
	}

	if len(sb) != len(st) {
		return false
	}

	for i := 0; i < len(sb); i++ {
		if bytes.IndexByte(sb, sb[i]) != bytes.IndexByte(st, st[i]) {
			return false
		}
	}

	return true
}

func change(s ...int) {
	s = append(s, 3)
}

type S struct {
	m string
}

func f() *S {
	return &S{"foo"} //A
}

func tes() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func f1() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func minTimeToVisitAllPoints(points [][]int) int {
	var sum int = 0
	for i := 0; i < len(points)-1; i++ {
		x := int(math.Abs(float64(points[i+1][0] - points[i][0])))
		y := int(math.Abs(float64(points[i+1][1] - points[i][1])))
		temp := x
		if y > x {
			temp = y
		}
		sum += temp
	}
	return sum
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {

	l := len(nums)
	for i := 0; i < l; i++ {
		if i+k < l {
			if abs(nums[i], nums[i+k]) == t {
				return true
			}
		}
	}

	return false
}

func abs(a, b int) int {
	if a-b > 0 {
		return a - b
	}

	return b - a
}

func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		flag := true
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
}

func selectSort(nums []int) {
	length := len(nums)
	for i := 0; i < length; i++ {
		maxIndex := 0
		//寻找最大的一个数，保存索引值
		for j := 1; j < length-i; j++ {
			if nums[j] > nums[maxIndex] {
				maxIndex = j
			}
		}
		nums[length-i-1], nums[maxIndex] = nums[maxIndex], nums[length-i-1]
	}
}

func fdsfds(nums []int) {
	l := len(nums)
	for i := 0; i < l; i++ {
		maxIndex := 0
		for j := 1; j < l-1; j++ {
			if nums[maxIndex] < nums[j] {
				maxIndex = j
			}
		}

		nums[l-i-1], nums[maxIndex] = nums[maxIndex], nums[l-i-1]
	}
}
