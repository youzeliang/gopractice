package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/cast"
)

func main() {

	// 创建一个 Redis 客户端实例
	client := redis.NewClient(&redis.Options{
		Addr:     "175.27.229.100:6380", // 指定 Redis 服务器地址和端口
		Password: "xxx",                 // 指定 Redis 服务器密码（如果有）
		DB:       0,                     // 使用默认数据库
	})

	for i := 0; i < 10000; i++ {
		_, err := client.LPush("QueueTitle", cast.ToString(i)).Result()
		if err != nil {
			fmt.Println(111)
		}
	}

}
