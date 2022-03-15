package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/spf13/cast"
)

func main() {

	//waitTime()
	//FF()
	Fortest()
}

func Fortest() {

	GoBirthTime := "2006-01-02 15:04:05"
	CreateTime := "2021-12-17 10:00:00"

	monthPlanMap := make(map[string][]int)
	stamp, _ := time.ParseInLocation(GoBirthTime, CreateTime, time.Local)
	fmt.Println(stamp.Unix())
	monthPlanMap[coverUnixToMonth(cast.ToInt64(stamp.Unix()))] = append(monthPlanMap[coverUnixToMonth(cast.ToInt64(stamp.Unix()))], 321)

	fmt.Println(monthPlanMap)

}

func coverUnixToMonth(startTime int64) string {
	// old table 2020/6/01
	if startTime > 1590940800 {
		return time.Unix(startTime, 0).Format("200601")
	}
	return "202005"
}

func FF() {

	for true {

		for i := 0; i < 100; i++ {

			defer fmt.Println(i)
		}
	}
}

func waitTime() {
	var w = sync.WaitGroup{}
	var ch = make(chan bool)
	w.Add(2)
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("我2秒")
		w.Done()
	}()
	go func() {
		time.Sleep(time.Second * 6)
		fmt.Println("我6秒")
		w.Done()
	}()
	go func() {
		w.Wait()
		ch <- false
	}()

	select {
	case <-time.After(time.Second * 5):
		fmt.Println("我超时了")
	case <-ch:
		fmt.Println("我结束了")
	}
}
