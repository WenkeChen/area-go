package cache

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/spf13/viper"
)

var RedisClient *redis.Client

func New() *redis.Client {
	if RedisClient == nil {
		RedisClient = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", viper.GetString("redis.host"), viper.GetString("redis.port")),
			Password: viper.GetString("reids.password"),
			DB:       0, // use default DB
		})

		pong, err := RedisClient.Ping().Result()
		if pong != "PONG" {
			panic(err)
		}
	}
	return RedisClient
}
