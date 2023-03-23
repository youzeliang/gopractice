package main

import (
	"fmt"
	"net/http"
)

func main() {
	for true {
		// case 1
		requestClose()
		// case 2
		//requestNotClose()
		//time.Sleep(time.Microsecond * 100)
	}
}

func requestNotClose() {
	_, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Printf("errInfo error: %s", err.Error())
	}
	fmt.Println("done")
}

func requestClose() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Printf("errInfo error:  %s", err.Error())
		return
	}
	defer resp.Body.Close()
	fmt.Println("ok")
}
