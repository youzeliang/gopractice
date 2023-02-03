package main

import (
	"fmt"
	"sort"
)

const (
	WinterActivitiesEndTime = "2022-12-19 20:00:00"
)

type GetUserMedalListReply_Item_MedalList struct {
	ActivityId int    `json:"activityId"`
	Name       string `json:"name"`
}

func main() {

	var a = GetUserMedalListReply_Item_MedalList{
		ActivityId: 2,
		Name:       "2222",
	}

	var b = GetUserMedalListReply_Item_MedalList{
		ActivityId: 1,
		Name:       "11",
	}

	var c = GetUserMedalListReply_Item_MedalList{
		ActivityId: 3,
		Name:       "333",
	}

	list := make([]GetUserMedalListReply_Item_MedalList, 0)

	list = append(list, a, b, c)

	sort.Sort(medalTtems(list))

	fmt.Println(list)

}

type medalTtems []GetUserMedalListReply_Item_MedalList

func (its medalTtems) Len() int { return len(its) }

func (its medalTtems) Less(i, j int) bool {
	return its[i].ActivityId < its[j].ActivityId
}

func (its medalTtems) Swap(i, j int) { its[i], its[j] = its[j], its[i] }
