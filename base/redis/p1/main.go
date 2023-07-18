package main

//import (
//	"fmt"
//	"github.com/go-redis/redis"
//)
//
//func main() {
//	client := redis.NewClient(&redis.Options{
//		Addr:     "175.27.229.100:6380",
//		Password: "xxxx",
//		DB:       0,
//	})
//
//	fmt.Println("----")
//	pong, _ := client.Ping().Result()
//	fmt.Println("----")
//	fmt.Println(pong)
//
//	client.Set("aaaa", "vvvv", -1)
//
//	//client.SetBit("key", 0, 1)
//	//client.SetBit("key", 1, 1)
//	//client.SetBit("key", 2, 0)
//	//client.SetBit("key", 3, 1)
//	//client.SetBit("key", 4, 1)
//
//	res := client.Get("aaaa").Val()
//	fmt.Println(res)
//}
