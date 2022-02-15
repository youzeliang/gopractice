package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"

	"github.com/sony/sonyflake"
)

type SortSliceTest struct {
	Num  int
	Name string
}

type GetStuAllDataListReply struct {
	List []GetStuAllData `json:"list"`
}

type GetStuAllData struct {
	PlayType    int   `json:"playType"`
	PadNum      int   `json:"padNum"`
	OnlineTime  int   `json:"onlineTime"`
	BOnlineTime int   `json:"bOnlineTime"`
	AOnlineTime int   `json:"aOnlineTime"`
	EnterTime   int64 `json:"enterTime"`
	FirstTime   int64 `json:"firstTime"`
	LastTime    int64 `json:"lastTime"`
	TDots       []int `json:"tDots"`
}

const aaaa = 1

func aaa() {

	if aaaa > -1 {
		aaaa = 3
	}

	fmt.Println(aaaa)
}

func main() {
	aaa()
	fmt.Println("-------------------------")

	var b = GetStuAllData{
		PlayType:    1,
		PadNum:      1,
		OnlineTime:  6600,
		BOnlineTime: 0,
		AOnlineTime: 0,
		EnterTime:   0,
		FirstTime:   0,
		LastTime:    0,
		TDots:       []int{6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6},
	}

	s := make([]GetStuAllData, 0)
	s = append(s, b)

	var a = GetStuAllDataListReply{
		List: s,
	}

	sdot := 0
	edot := 0
	onlineTimeMax := 0
	for inx, dinfo := range a.List {
		if a.List[inx].OnlineTime > onlineTimeMax {
			onlineTimeMax = a.List[inx].OnlineTime
		}
		// 找第一个主讲在线打点的心跳sx和最后一个主讲在线的心跳ex
		sx := 0
		ex := len(dinfo.TDots) - 1
		for ; sx < len(dinfo.TDots); sx++ {
			if dinfo.TDots[sx] == 6 {
				break
			}
		}
		for ; ex >= 0; ex-- {
			if dinfo.TDots[ex] == 6 {
				break
			}
		}

		// 如果是直播时长，需要对齐进入的偏移时间
		if dinfo.PlayType == 1 {
			sx += dinfo.PadNum
			ex += dinfo.PadNum
			if sx < 0 {
				sx = 0
			}
		}

		if inx == 0 {
			sdot = sx
			edot = ex
		} else {
			// 做并集
			if sx < sdot {
				sdot = sx
			}
			if ex > edot {
				edot = ex
			}
		}
	}

	fmt.Println(edot)
	fmt.Println(sdot)

	if edot < sdot {
		fmt.Println(111)
	}
	fmt.Println(onlineTimeMax)
	if onlineTimeMax < (edot - sdot + 1) {
		fmt.Println(222)
	}

}

func InitArrs() (arrs []SortSliceTest) {
	//arr1 := SortSliceTest{
	//	Num:  3,
	//	Name: "arr1",
	//}
	//
	//arr2 := SortSliceTest{
	//	Num:  1,
	//	Name: "arr2",
	//}
	//
	//arr3 := SortSliceTest{
	//	Num:  5,
	//	Name: "arr3",
	//}
	//
	//arr4 := SortSliceTest{
	//	Num:  2,
	//	Name: "arr4",
	//}
	//
	//arrs = append(arrs, arr1, arr2, arr3, arr4)
	return
}

var flake = sonyflake.NewSonyflake(sonyflake.Settings{})

func GetUuidString() string {
	return strconv.FormatUint(GetUuidUInt64(), 10)
}

func GetUuidUInt64() uint64 {
	uid, _ := flake.NextID()
	return uid
}

func PerformRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}

	return data
}

func GetGreeting(name string) (greeting string) {
	greeting = "Hello " + name + "."
	return
}
