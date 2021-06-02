package main

import "fmt"

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



	fmt.Println(TDOTTYPE[0][true])



}
