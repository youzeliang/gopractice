package main

func main() {

	fillOne(1000000, 1000000)
}

func fillOne(m, n int) {

	res := [1000][1000]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[j][i] = 0
		}
	}
}

func main1() {

	a := make([]int, 0)

	buffer := make(chan int)

	go func() {
		for v := range buffer {
			a = append(a, v)
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			buffer <- i
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println(a)

}
