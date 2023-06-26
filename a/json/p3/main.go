package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"strconv"
)

var Json = jsoniter.ConfigCompatibleWithStandardLibrary

type Bird struct {
	Species     string
	Description string
}

type PercentageList struct {
	Percentage float64 `json:"percentage"`
	OnStatus   int     `json:"onStatus"`
}

var (
	OnlineStatus map[int]int
	Online       map[int]struct{}
)

const (
	Offline = iota
	OnlineTutor
	OfflineTutor
	ClassOffline = iota + 1
	ClassOnline
	MainOnline
	MainOffline
)

const (
	MOffline = iota
	MOnline
	MHangUp
)

func init() {
	OnlineStatus = map[int]int{
		OnlineTutor:  MOnline,
		ClassOnline:  MOnline,
		MainOnline:   MOnline,
		Offline:      MOffline,
		OfflineTutor: MHangUp,
		ClassOffline: MHangUp,
		MainOffline:  MHangUp,
	}

	Online = map[int]struct{}{
		OnlineTutor: {},
		ClassOnline: {},
		MainOnline:  {},
	}

}

func main() {

	var TDots = []int{1, 1, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6}
	TDots = append(TDots, 0)
	l := len(TDots)

	list := make([]PercentageList, 0)
	TDots = append(TDots, 0)
	count := 1
	//var ontime int
	for m := 1; m < l; m++ {
		if TDots[m] != TDots[m-1] {
			fmt.Println(TDots[m])
			value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(count)/float64(l-1)), 64)
			list = append(list, PercentageList{
				Percentage: value * 100,
				OnStatus:   OnlineStatus[TDots[m-1]],
			})
			count = 1
			continue
		}
		count++

		//if _, ok := Online[TDots[m]]; ok {
		//	ontime++
		//}

	}

	fmt.Println(list)

}

//if i >= recordNum {
//break
//}
//onlineTime := 0
//pList := make([]quality.PercentageList, 0)
//l := len(j.TDots)
//if l == 0 {
//continue
//}
//count := 1
//j.TDots = append(j.TDots, 0)
//
//if j.PlayType != thirdApi.PlayTypeBack {
//if j.PadNum > 0 {
//var left []int
//for x := 0; x < j.PadNum; x++ {
//left = append(left, 0)
//}
//j.TDots = append(left, j.TDots...)
//} else if j.PadNum < 0 {
//j.TDots = j.TDots[-j.PadNum:]
//}
//if len(j.TDots) < total {
//for x := 0; x < total-len(j.TDots); x++ {
//j.TDots = append(j.TDots, 0)
//}
//}
//} else {
//if total > l {
//for n := 0; n < total-l; n++ {
//j.TDots = append(j.TDots, 0)
//}
//}
//}
//j.TDots = append(j.TDots, 10)
//l = len(j.TDots)
//if _, ok := thirdApi.OnlineBack[j.TDots[0]]; ok {
//onlineTime++
//}
//for m := 1; m < l; m++ {
//if j.PlayType == thirdApi.PlayTypeBack {
//if _, ok := thirdApi.OnlineBack[j.TDots[m]]; ok {
//onlineTime++
//}
//}
//
//if j.TDots[m] != j.TDots[m-1] {
//
//value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(count)/float64(l-1)), 64)
//status := thirdApi.OnlineStatus[j.TDots[m-1]]
//if j.PlayType == thirdApi.PlayTypeBack && status == thirdApi.PlayTypeBack {
//status = thirdApi.MOnline
//}
//pList = append(pList, quality.PercentageList{
//Percentage: value * common.PercentageRate,
//OnStatus:   status,
//})
//count = 1
//continue
//}
//count++
//}
//
//if j.PlayType != thirdApi.PlayTypeBack {
//var p float64
//for x := 0; x < len(pList); x++ {
//if pList[x].OnStatus == thirdApi.MOnline {
//p += pList[x].Percentage
//}
//}
//
//onlineTime = total * int(p) / common.PercentageRate
//}
//
//planTime := int((baseInfo.LiveEnd - baseInfo.LiveStart) / common.TimeRate)
//if onlineTime > planTime {
//onlineTime = planTime
//}
//
//if onlineTimeMax < onlineTime {
//onlineTimeMax = onlineTime
//}
//
//temp := float64(0)
//for f := 0; f < len(pList); f++ {
//if f == len(pList)-1 && len(pList)-1 >= 0 {
//if common.PercentageRate-temp > pList[len(pList)-1].Percentage {
//pList[len(pList)-1].Percentage = common.PercentageRate - temp
//}
//break
//}
//temp += pList[f].Percentage
//}
//
//dList = append(dList, quality.OnlineTimeInfo{
//PlayType:   j.PlayType,
//OnlineTime: onlineTime,
//List:       pList,
//})
//}
//
//baseInfo.StuOnlineTime = onlineTimeMax * common.TimeRate
