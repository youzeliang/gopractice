package main

import (
	"time"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:         "10.170.124.79:6379", // conf.Conf.Redis.Addr,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     100,
		PoolTimeout:  30 * time.Second,
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {

	}
}

func main() {
	InitRedis()
}
