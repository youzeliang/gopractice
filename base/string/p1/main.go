package main

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