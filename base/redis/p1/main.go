package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	pong, _ := client.Ping().Result()
	fmt.Println(pong)

	client.Set("aaaa", "vvvv", -1)

	client.SetBit("key", 0, 1)
	client.SetBit("key", 1, 1)
	client.SetBit("key", 2, 0)
	client.SetBit("key", 3, 1)
	client.SetBit("key", 4, 1)

	res := client.Get("bmk").Val()
	fmt.Println(res)
}
