package main

<<<<<<< HEAD
import "fmt"

func main() {

	var a = []int{1, 1, 2, 3, 4, 5}

	cnts := map[int]int{}

	for i := 0; i < len(a); i++ {
		cnts[a[i]]++
	}

	fmt.Println(cnts)
=======
import (
	"sync"
)

func main() {

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
>>>>>>> 9ffb43e3806758d7cdbde180b68cce6f43591c8d
}
