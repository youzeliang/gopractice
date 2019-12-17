package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserId   int
	UserName string
}
type UserTag struct {
	UserId   int    `json:"user_id" bson:"user_id"`
	UserName string `json:"user_name" bson:"user_name"`
}

func main() {
	u := &User{UserId: 1, UserName: "Murphy"}
	j, _ := json.Marshal(u)
	fmt.Printf("struct User echo : %v\n", string(j))

	u_tag := &UserTag{UserId: 1, UserName: "Murphy"}
	j_tag, _ := json.Marshal(u_tag)
	fmt.Printf("struct UserTag echo : %v\n", string(j_tag))

}
