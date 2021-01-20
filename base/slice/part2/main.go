package main

func main() {
	//num := 6
	//for index := 0; index < num; index++ {
	//	resp, _ := http.Get("https://www.baidu.com")
	//	_, _ = ioutil.ReadAll(resp.Body)
	//}
	//fmt.Printf("此时goroutine个数= %d\n", runtime.NumGoroutine())

	c()

}

func c() {

	var cc chan int

	cc <- 1

}
