package main

import (
	"fmt"
	"net/http"
	"time"
)

type S int

func main() {
	//http.HandleFunc("/", SayHello) // 设置访问的路由
	//
	//log.Fatalln(http.ListenAndServe(":8080", nil))

	a := S(0)
	b := make([]*S, 2)
	b[0] = &a
	c := new(S)
	b[1] = c

}

func SayHello(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(&request)

	go func() {
		for range time.Tick(time.Second) {
			select {
			case <-request.Context().Done():
				fmt.Println(11)
				return
			}
		}
	}()

	time.Sleep(2 * time.Second)
	writer.Write([]byte("Hi"))
}
