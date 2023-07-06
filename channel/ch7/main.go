package main

//import (
//	"fmt"
//)
//
//func main() {
//
//<<<<<<< HEAD
//	ch := prodece(1, 2, 3, 4)
//	otu := square(ch)
//	for n := range otu {
//		fmt.Println(n)
//	}
//
//}
//
//func prodece(nums ...int) <-chan int {
//	out := make(chan int)
//	go func() {
//		defer close(out)
//		for _, n := range nums {
//			out <- n
//		}
//	}()
//
//	return out
//}
//
//func square(inCh <-chan int) <-chan int {
//
//	out := make(chan int)
//
//	go func() {
//		defer close(out)
//		for n := range inCh {
//			out <- n * n
//		}
//	}()
//
//	return out
//=======
//	cacheCh := make(chan int, 5)
//	cacheCh <- 2
//	cacheCh <- 3
//	fmt.Println("cacheCh容量为:", cap(cacheCh), ",元素个数为：", len(cacheCh))
//}
