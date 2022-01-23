package database

import (
	"github.com/go-redis/redis/v8"
)

func SetupRedis(addr string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   0,
	})
}
