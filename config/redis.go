package config

import (
	"github.com/go-redis/redis"
	"log"
	"newgo/global"
)

func InitRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     AppConfig.Redis.Host + AppConfig.Redis.Port,
		Password: AppConfig.Redis.Password,
		DB:       AppConfig.Redis.DB,
	})
	if _, err := RedisClient.Ping().Result(); err != nil {
		log.Fatalf("connect redis failed, err:%v", err)
	}
	global.RedisClient = RedisClient
}
