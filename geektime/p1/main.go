package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(2314)
	})

	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalln(err)
		}

	}()

	select {}
}
