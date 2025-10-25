package cache

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func ConnectToRedis() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb, nil
}
