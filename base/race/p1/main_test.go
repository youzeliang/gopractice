package main

import (
	"sync"
	"testing"
)

func TestName(t *testing.T) {

	//x := []string{"321"}
	x := make([]string, 0, 6)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		y := append(x, "hell0", "11")
		t.Log(cap(y), len(y))
	}()

	go func() {
		defer wg.Done()
		z := append(x, "world", "22")
		t.Log(cap(z), len(z))
	}()

	wg.Wait()
}
