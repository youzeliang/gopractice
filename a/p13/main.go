package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var a = map[string]string{
	"11": "111",
	"22": "22",
}

func main() {

	//r := gin.Default()
	//
	//r.GET("/", func(c *gin.Context) {
	//	c.String(http.StatusOK, "hello World!")
	//})
	//// 3.监听端口，默认在8080
	//// Run("里面不指定端口号默认为8080")
	//r.Run(":8000")

	//for i := 0; i < 100; i++ {
	//	fmt.Println(getRand(10))
	//}

	//m :=make(map[string]float64,0)
	//
	//fmt.Println( m["fff"])

	timeStr := time.Now().Format("2006-01-02")
	//fmt.Println("timeStr:", timeStr)
	t, _ := time.Parse("2006-01-02", timeStr)
	fmt.Println(time.Now().Unix())

	fmt.Println(t.Add(-8 * time.Hour).Unix())

}

func getRand1(num int) int {
	rand.Seed(time.Now().UnixNano())
	v := rand.Intn(num)
	return v
}

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

func getRand(num int) int {
	rand.Seed(time.Now().UnixNano())
	var mu sync.Mutex
	mu.Lock()
	v := rand.Intn(num)
	mu.Unlock()
	return v
}

type A struct {
	m map[string]string
}

func (a *A) Init() {

	a.m = make(map[string]string, 0)
}

func (a *A) get(s string) string {
	return a.m[s]
}

func (a *A) set(k, v string) {
	a.m[k] = v
}
