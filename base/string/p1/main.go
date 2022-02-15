package main

<<<<<<< HEAD
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
>>>>>>> 019d9790d25543eb092bf612e47a058923e13fff
