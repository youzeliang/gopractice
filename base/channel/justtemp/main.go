package main

import (
	"fmt"
	"github.com/spf13/cast"
	"strconv"
	"time"
)

type durationInfo struct {
	TDots    []int
	PlayType int
	PadNum   int
}

type PercentageList struct {
	Percentage float64 `json:"percentage"`
	OnStatus   int     `json:"onStatus"`
}

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

const PlayTypeBack = 2

var (
	OnlineStatus map[int]int
	Online       map[int]struct{}
	OnlineBack   map[int]struct{}
)

func init() {
	OnlineStatus = map[int]int{
		OnlineTutor:  MOnline,
		ClassOnline:  MOnline,
		MainOnline:   MOnline,
		Offline:      MOffline,
		OfflineTutor: MHangUp,
		ClassOffline: MHangUp,
		MainOffline:  MOnline,
	}

	Online = map[int]struct{}{
		OnlineTutor: {},
		ClassOnline: {},
		MainOnline:  {},
	}

	OnlineBack = map[int]struct{}{
		OnlineTutor: {},
		ClassOnline: {},
		MainOnline:  {},
		MainOffline: {},
	}

}

func main() {

	aa := "2022-04-08 14:24:09"

	xx := time.Unix(cast.ToInt64(aa), 0).Format("2006-01-02 15:04:05")
	fmt.Println(xx)

	total := 90
	var db = []int{1, 1, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6}

	co := 0
	for i := 0; i < len(db); i++ {
		if db[i] != 0 {
			co++
		}
	}

	fmt.Println("--------",co)

	var a = durationInfo{
		TDots:    db,
		PlayType: 1,
		PadNum:   -1,
	}
	list := make([]durationInfo, 0)
	list = append(list, a)

	pList := make([]PercentageList, 0)

	for i, j := range list {
		if i >= 2 {
			break
		}
		onlineTime := 0
		l := len(j.TDots)
		if l == 0 {
			continue
		}
		count := 1
		j.TDots = append(j.TDots, 0)

		if j.PlayType != 2 {
			if j.PadNum > 0 {
				var left []int
				for x := 0; x < j.PadNum; x++ {
					left = append(left, 0)
				}
				j.TDots = append(left, j.TDots...)
			} else if j.PadNum < 0 && -j.PadNum < len(j.TDots) {
				j.TDots = j.TDots[-j.PadNum:]
			}
			if len(j.TDots) < total {
				for x := 0; x < total-len(j.TDots); x++ {
					j.TDots = append(j.TDots, 0)
				}
			} else if len(j.TDots) > total {
				j.TDots = j.TDots[:total]
			}

			fmt.Println(j.TDots)

		} else {
			if total > l {
				for n := 0; n < total-l; n++ {
					j.TDots = append(j.TDots, 0)
				}
			}
		}
		j.TDots = append(j.TDots, 10)
		x := 0
		for i := 0; i < len(j.TDots); i++ {
			if j.TDots[i] == 6 {
				x++
			}
		}

		l = len(j.TDots)

		if _, ok := OnlineBack[j.TDots[0]]; ok {
			onlineTime++
		}

		fmt.Println(j.TDots)

		for m := 1; m < l; m++ {
			if j.PlayType == PlayTypeBack {
				if _, ok := OnlineBack[j.TDots[m]]; ok {
					onlineTime++
				}
			}

			if j.TDots[m] != j.TDots[m-1] {

				value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(count)/float64(l-1)), 64)
				status := OnlineStatus[j.TDots[m-1]]
				if j.PlayType == PlayTypeBack && status == PlayTypeBack {
					status = MOnline
				}
				pList = append(pList, PercentageList{
					Percentage: value * 100,
					OnStatus:   status,
				})
				count = 1
				continue
			}
			count++
		}

		if j.PlayType != PlayTypeBack {
			var p float64
			for x := 0; x < len(pList); x++ {
				if pList[x].OnStatus == MOnline {
					p += pList[x].Percentage
				}
			}

			onlineTime = total * int(p) / 100
		}

		temp := float64(0)
		for f := 0; f < len(pList); f++ {
			if f == len(pList)-1 && len(pList)-1 >= 0 {
				if 100-temp > pList[len(pList)-1].Percentage {
					pList[len(pList)-1].Percentage = 100 - temp
				}
				break
			}
			temp += pList[f].Percentage
		}

		fmt.Println("////", onlineTime)

	}

	fmt.Println(pList)

}
