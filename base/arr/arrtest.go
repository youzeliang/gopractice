package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
)

const name = 7

func main() {
	//Ts()

	ffaa()
}

func ffaa() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Geektutu")
	})
	r.Run() // listen and serve on 0.0.0.0:8080

}

func Ts() {

	slice := []int{1, 2}
	//slice := array[0:2]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	newSlice := append(slice, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13)
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	rand.Int()

	//m := make(map[string]string, 0)

	var apr = struct {
		Name string
		Age  int
	}{
		Name: "fdsfds",
		Age:  10,
	}

	fmt.Println(apr)

}
