package db

import (
	"context"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func RedisConnection() (*redis.Client, context.Context) {
	dbNum, err := strconv.Atoi(os.Getenv("REDIS_DATABASE"))
	if err != nil {
		dbNum = 0
	}
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       dbNum,
	})

	return rdb, context.Background()
}
