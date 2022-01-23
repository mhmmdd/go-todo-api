package utils

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
	"time"
)

func BoolAddr(b bool) *bool {
	boolVar := b
	return &boolVar
}

var CacheChannel chan string

func SetupCacheChannel(r *redis.Client) {
	CacheChannel = make(chan string)

	go func(ch chan string) {
		for {
			time.Sleep(3 * time.Second)

			key := <-ch
			r.Del(context.Background(), key)
			fmt.Printf("Cache cleared: %s\n", key)
		}
	}(CacheChannel)
}

func ClearCache(keys ...string) {
	for _, key := range keys {
		CacheChannel <- key // channel receives the key
	}
}
