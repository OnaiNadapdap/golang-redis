package database

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/onainadapdap1/golang-crud-redis/config"
)

func ConnectionRedisDB(config *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: config.RedisUrl,
	})
	fmt.Println("connected successfully to the database (redis)")
	return rdb
}