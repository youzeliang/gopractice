package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"time"
)

func main() {

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	go func() {
		for {
			fmt.Println(<-ticker.C)
		}
	}()

	time.Sleep(time.Second * 10)
	fmt.Println("finish")

}

type StuInfo struct {
	StuId    int `json:"stuId"`
	StuCouId int `json:"stuCouId"`
}

func GetSign(stuId, stuCouId, planId string) (token string) {
	sign := "Xes-Report-Token-" + stuId + "-" + stuCouId + "-" + planId
	return Md5(sign)
}

func Md5(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}
