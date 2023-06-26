package main

<<<<<<< HEAD
import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4}
	newSlice := slice[0:3]
	fmt.Println("before modifying underlying array:")
	fmt.Println("slice: ", slice)
	fmt.Println("newSlice: ", newSlice)
	fmt.Println()

	newSlice[0] = 6
	newSlice = append(newSlice, 1, 2, 3, 4, 5, 6)
	fmt.Println("after modifying underlying array:")
	fmt.Println("slice: ", slice)
	fmt.Println("newSlice: ", newSlice)
=======
import "fmt"

func main() {
	foo := []int{0, 0, 0, 42, 100}
	bar := foo[1:4]
	fmt.Println(foo)
	bar[1] = 99
	fmt.Println("foo:", foo)
	fmt.Println("bar:", bar)
>>>>>>> 9dad412f07f955fe7cf7c2792221d5e58f71b91e
}
