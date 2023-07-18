package main

//import (
//	"fmt"
//	"github.com/go-redis/redis"
//)
//
//func main() {
//	// 创建一个 Redis 客户端实例
//	client := redis.NewClient(&redis.Options{
//		Addr:     "175.27.229.100:6380", // 指定 Redis 服务器地址和端口
//		Password: "Lll271828",           // 指定 Redis 服务器密码（如果有）
//		DB:       0,                     // 使用默认数据库
//	})
//
//	// 订阅一个频道
//	pubsub := client.Subscribe("mychannel")
//	defer pubsub.Close()
//
//	// 接收频道消息的回调函数
//	go func() {
//		for msg := range pubsub.Channel() {
//			fmt.Println(msg.Channel, msg.Payload)
//		}
//	}()
//
//	// 发布一条消息到频道
//	err := client.Publish("mychannel", "Hello, world!").Err()
//	if err != nil {
//		panic(err)
//	}
//
//	// 等待接收到一条消息
//	select {}
//}
