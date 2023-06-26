package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	Srv *UserCenter
)

func Init() {
	Srv = new(UserCenter)
}

func New() *UserCenter {
	return &UserCenter{
		studentService: NewServiceS(Srv.studentService),
		teacherService: NewServictt(Srv.teacherService),
	}
}

func main() {

	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.Parse("2006-01-02", timeStr)
	zeroPoint := t.AddDate(0, 0, 1).Unix()

	fmt.Println(zeroPoint)

	var MissDuration = "6,6,6,6,6,6"

	durationList := strings.Split(MissDuration, ",")
	var durationIntList []int
	//var onlineTime int
	for i := 0; i < len(durationList); i++ {
		a, _ := strconv.Atoi(durationList[i])
		durationIntList = append(durationIntList, a)
		//fmt.Println(durationList[i])
	}

	struts, err := json.Marshal(durationIntList)
	if err != nil {

		fmt.Println(err)

	}

	fmt.Println(string(struts))
	//
	//fmt.Println(len(a))

}

type UserCenter struct {
	studentService Ss
	teacherService St
	onceLock       sync.Once
}

type Ss interface {
	Print(a int)
}

type St interface {
	Print(a int)
}

type SSss struct {
	sssss Ss
}

func NewServiceS(s Ss) *SSss {
	return &SSss{
		sssss: s,
	}
}

func NewServictt(s St) *SStt {
	return &SStt{
		ssssttt: s,
	}
}

type SStt struct {
	ssssttt St
}

func (s *SSss) Print(a int) {
	fmt.Println(a)
}

func (s *SStt) Print(a int) {
	fmt.Println(a)
}

func (t *UserCenter) AA() {

	New().teacherService.Print(11)

}
