package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {

	// 创建一个 Redis 客户端实例
	client := redis.NewClient(&redis.Options{
		Addr:     "175.27.229.100:6380", // 指定 Redis 服务器地址和端口
		Password: "xxxx",                // 指定 Redis 服务器密码（如果有）
		DB:       0,                     // 使用默认数据库
	})

	for {
		res, err := client.BRPop(0, "QueueTitle").Result()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(res[1])
	}

}
