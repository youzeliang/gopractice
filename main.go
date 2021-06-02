package main

import (
	"encoding/json"
	"fmt"
)

var (
	//辅导在线     = 1
	TUTOR_ONLINE = 1
	//辅导挂机     = 2
	TUTOR_SUSPEND = 2
	//主讲在线     = 6
	MAINTEACHER_ONLINE = 6
	//主讲挂机     = 7
	MAINTEACHER_SUSPEND = 7
	BACKEND_ONLINE      = 8
	//主讲态      = uint8(1)
	PLANSTATUS_MAINTEACHER = uint8(1)
	//辅导态      = uint8(2)
	PLANSTATUS_TUTOR = uint8(2)
	//课前辅导     = uint8(3)
	BEFORECLASS_TUTOR = uint8(3)
	//课后辅导     = uint8(4)
	AFTERCLASS_TUTOR = uint8(4)
	//是挂机      = true
	SUSPEND = true
	//否挂机      = false
	NOTSUSPEND = false
	TDOTTYPE   = map[uint8]map[bool]int{
		PLANSTATUS_TUTOR:       {NOTSUSPEND: TUTOR_ONLINE, SUSPEND: TUTOR_SUSPEND},
		BEFORECLASS_TUTOR:      {NOTSUSPEND: TUTOR_ONLINE, SUSPEND: TUTOR_SUSPEND},
		AFTERCLASS_TUTOR:       {NOTSUSPEND: TUTOR_ONLINE, SUSPEND: TUTOR_SUSPEND},
		PLANSTATUS_MAINTEACHER: {NOTSUSPEND: MAINTEACHER_ONLINE, SUSPEND: MAINTEACHER_SUSPEND},
	}
)

func main() {


	if _, ok := TDOTTYPE[0][true]; ok != true {

		fmt.Println(321321)
	}

	type name struct {
		All   float64 `json:"all"`
		Redis float64 `json:"redis"`
		Mysql float64 `json:"mysql"`
	}

	s := "{\"all\":0.01,\"redis\":0.4,\"mysql\":0.01,\"rpc\":0.01,\"http\":0.01,\"kafka\":0.01,\"etcd\":0.01}"

	var revMsg name
	err := json.Unmarshal([]byte(s), &revMsg)
	if err != nil {
		fmt.Println(err)
	}


	ss := `{"all":["x_param","x_response"],"talcamp":["x_param","x_response"]}`


	var m map[string][]string

	err = json.Unmarshal([]byte(ss), &m)
	if err != nil {
		fmt.Println(err)
	}

	for i,j := range m{
		fmt.Println(i)
		fmt.Println(j)
	}


	fmt.Println(revMsg.Redis)

}
