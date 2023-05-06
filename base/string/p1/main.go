package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
		s()
}

func s() {

	defer func() {
		if err := recover(); err != nil {
			go a()
		}
	}()
	a()

	select {

	}

}

func a() {
	ss := rand.Intn(1)

	fmt.Println(ss, "--------")
	if 2 == 2 {
		fmt.Println(331)
		panic("ffff")
	}

	fmt.Println("end")
}

import "fmt"

func main()  {
	str := "Golang罗尔"
	fmt.Println(len(str))
	fmt.Println(len([]rune(str)))

	PlanStatus := 6
	if PlanStatus > 5 || PlanStatus <= 0 {
		fmt.Println(111)
	}


}

//timeHit.PlanStatus > cast.ToUint8(define.PlanStatusMax) || timeHit.PlanStatus <= 0
=======
import (
	"fmt"
	"strings"
)

func main() {

	s := "fdsfsd"

	isExist := strings.Index(s[0:0], string(s[0]))

	fmt.Println(isExist)

}

func len(nums []int) int {
	return 0
}