package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

func main() {
	// 开启pprof
	go func() {
		ip := "0.0.0.0:6060"
		if err := http.ListenAndServe(ip, nil); err != nil {
			fmt.Printf("start pprof failed on %s\n", ip)
			os.Exit(1)
		}
	}()

	tick := time.Tick(time.Second * 1)
	var buf []byte
	stime := time.Now()

	for range tick {
		// 1秒1M内存
		buf = append(buf, make([]byte, 1024*1024)...)

		fmt.Printf("%f\n", time.Now().Sub(stime).Seconds())
	}
}
