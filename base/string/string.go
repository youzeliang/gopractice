package main

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

func main() {
	//changeString()

	//AAA()

	//args.StudentId < 0 && (args.BizId != 2 && args.BizId != 3


	//jsutil.Json.UnmarshalFromString

}

func AAA() {

	str := stringToStract()
	var a []name

	err := json.Unmarshal([]byte(str), &a)
	if err != nil {

	}

	for i := 0; i < len(a); i++ {

		fmt.Println(a[i])
	}

}

type name struct {
	StuCouId  int    `json:"stuCouId"`
	Duration  string `json:"duration"`
	LastTime  int    `json:"lastTime"`
	FirstTime int    `json:"firstTime"`
}

type nameA struct {
	StuCouId  int    `json:"stuCouId"`
	Duration  string `json:"duration"`
	LastTime  int    `json:"lastTime"`
	FirstTime int    `json:"firstTime"`
}

func stringToStract() string {

	var list []name

	for i := 1; i < 10; i++ {

		var a = name{
			StuCouId:  i,
			Duration:  "",
			LastTime:  i,
			FirstTime: i,
		}
		list = append(list, a)
	}

	res, _ := json.Marshal(list)
	//fmt.Println(string(res))

	return string(res)

}

func changeString() {

	s1 := "hello"

	byte1 := []byte(s1)

	byte1[0] = 'e'

	fmt.Println(string(byte1))

	var a int32 = 43342

	println(unsafe.Sizeof(a))

}
