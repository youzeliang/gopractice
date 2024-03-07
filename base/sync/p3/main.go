package main

import (
	"bytes"
	"fmt"
	"sync"
	"time"
)

var pool *sync.Pool

func init() {
	pool = &sync.Pool{
		New: func() interface{} {
			return &bytes.Buffer{}
		},
	}
}

func withoutPool() {

	start := time.Now()

	for i := 0; i < 100000000; i++ {
		buf := &bytes.Buffer{}
		buf.WriteString("hello")
		buf.WriteString("world")
	}

	fmt.Println("Without pool:", time.Since(start))

}

func withPool() {
	start := time.Now()

	for i := 0; i < 100000000; i++ {

		buf := pool.Get().(*bytes.Buffer)
		buf.WriteString("hello")
		buf.WriteString("world")
		pool.Put(buf)
	}

	fmt.Println("With pool:", time.Since(start))
}

func main() {

	withPool()
	withoutPool()
}
