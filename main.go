package main

func main() {

	a := [2][2]int{
		{0, 1}, /*  第一行索引为 0 */
		{4, 5}, /*  第二行索引为 1 */
	}

	originMap := make(map[int]int)

	for _, j := range a {
		for n := range j {

			originMap[j[n]-j[n-1]] = j[n] - j[n-1]
			if _, ok := originMap[j[n]-j[n-1]]; ok {

			}

			//fmt.Println("11",n)

		}

		//fmt.Println(i)
		//fmt.Println(j)
	}
}
