package main

import "fmt"

var (
	firstName, lastName int
	i int
	f float32
	input = "56.12 / 5212 / Go"
	format = "%f / %d / %s"
)

func main() {
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	fmt.Println(firstName)
	//// fmt.Scanf("%s %s", &firstName, &lastName)
	//fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
	//fmt.Sscanf(input, format, &f, &i, &s)
	//fmt.Println("From the string we read: ", f, i, s)
	// 输出结果: From the string we read: 56.12 5212 Go
}

