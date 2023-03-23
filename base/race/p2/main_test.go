package main

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {

	fmt.Println(getNumber())
}

func getNumber() int {
	var i int
	go func() {
		i = 5
	}()

	return i
}
