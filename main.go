package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
	"unsafe"
)

type F struct {
	Name string
}

func NewF(name string) *F {
	return &F{Name: name}
}

func main() {

	go func() {
		ip := "0.0.0.0:6060"
		if err := http.ListenAndServe(ip, nil); err != nil {
			fmt.Printf("start pprof failed on %s\n", ip)
			os.Exit(1)
		}
	}()

	tick := time.Tick(time.Second / 100)
	var buf []byte
	for range tick {
		buf = append(buf, make([]byte, 1024*1024)...)
	}

}

func Float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}


