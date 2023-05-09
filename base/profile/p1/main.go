package main

import (
	"net/http"
	_ "net/http/pprof"
)

func main() {
	err := http.ListenAndServe("127.0.0.1:10001", nil)
	if err != nil {
		return
	}
}
