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

	var TDots = []int{1, 2, 2}
	l := len(TDots)

	list := make([]PercentageList, 0)
	count := 1
	var ontime int
	for m := 1; m < l+1; m++ {

		if TDots[m] != TDots[m-1] {
			value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(count)/float64(l)), 64)
			list = append(list, PercentageList{
				Percentage: value * 100,
				OnStatus:   OnlineStatus[TDots[m-1]],
			})
			count = 1
		}
		count++

		if _, ok := Online[TDots[m]]; ok {
			ontime++
		}

	}

	fmt.Println(list)

}
