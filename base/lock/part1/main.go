package main

//
//import (
//	"fmt"
//	"time"
//)
//
////func main() {
////	st := time.Now()
////	ch := make(chan bool, 2)
////	go func ()  {
////		time.Sleep(time.Second * 2)
////		<-ch
////	}()
////	ch <- true
////	ch <- true // 缓冲区为 2，发送方不阻塞，继续往下执行
////	fmt.Printf("cost %.1f s\n", time.Now().Sub(st).Seconds()) // cost 0.0 s
////	ch <- true // 缓冲区使用完，发送方阻塞，2s 后接收方接收到数据，释放一个插槽，继续往下执行
////	fmt.Printf("cost %.1f s\n", time.Now().Sub(st).Seconds()) // cost 2.0 s
////	time.Sleep(time.Second * 5)
////}
//
//func f1() {
//	for i := 0; i < 5; i++ {
//		defer fmt.Println(i)
//	}
//}
//
//func f2() {
//	for i := 0; i < 6; i++ {
//		defer func() {
//			fmt.Println(i)
//		}()
//	}
//}
//
//func f3() {
//	for i := 0; i < 5; i++ {
//		defer func(n int) {
//			fmt.Println(n)
//		}(i)
//	}
//}
//
//func f4() int {
//	t := 5
//	defer func() {
//		t++
//	}()
//	return t
//}
//
//func f5() (r int) {
//	defer func() {
//		r++
//	}()
//	return 0
//}
//
//func f6() (r int) {
//	t := 5
//	defer func() {
//		t = t + 5
//	}()
//	return t
//}
//
//func f7() (r int) {
//	defer func(r int) {
//		r = r + 5
//	}(r)
//	return 1
//}
//
//func a()  int{
//	i := 0
//	defer fmt.Println(i)
//	i++
//	return i
//}
//=======
//	"sync"
//)
//
//var counter int
//
//func main() {
//
//	var wg sync.WaitGroup
//	var l sync.Mutex
//
//	for i := 0; i < 100; i++ {
//		wg.Add(1)
//		go func() {
//			defer wg.Done()
//			l.Lock()
//			counter++
//			l.Unlock()
//		}()
//	}
//
//	wg.Wait()
//
//	fmt.Println(counter)
//}
