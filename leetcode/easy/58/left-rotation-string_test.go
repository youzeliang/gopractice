package _58

import (
	"sort"
	"testing"
)

func Test_reverseLeftWords(t *testing.T) {
	//
	//s :="abc"
	//
	//for i := range s{
	//	fmt.Println(s[i])
	//}

	var a = [][]int{
		{1, 2}, /*  第一行索引为 0 */
		{2, 3}, /*  第二行索引为 1 */
		{3, 4}, /*  第二行索引为 1 */

	}

	//maxEvents(a)
	c := make([]int,0)

	for _,j  := range a {
		for m,n := range j{
			if m ==1{
				c = append(c, n)
			}
		}
	}

	sort.Ints(c)
	for _,j := range c{
		println(j)
	}
	}

func maxEvents(events [][]int) int {

	a := make([]int,0)
	b :=make([]int,0)

	for _,j  := range events {
		for m,n := range j{
			if m ==0{
				a = append(a, n)
			}else {
				b = append(b, n)
			}
		}
	}



	return min(a[len(a)-1]-a[0]+1,len(events))

}

func min(a,b int)int  {
	if a > b {
		return b
	}
	return a
}