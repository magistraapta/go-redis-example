package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func NewClientRedis() *redis.Client {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: 0, Protocol: 3})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Redis")

	return rdb
}
