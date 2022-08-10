package redis

import (
	"leaf_chat/conf"
	"log"

	"github.com/go-redis/redis"
)

var Client *redis.Client

func Init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     conf.Server.RedisAddr,
		Password: conf.Server.RedisPassword,
		DB:       conf.Server.RedisDb,
	})

	if _, err := Client.Ping().Result(); err != nil {
		log.Fatalf("redis connect err : %s", err.Error())
	}
}

func Close() {
	defer Client.Close()
}
