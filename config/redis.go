package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var RD *redis.Client

func InitRedis() {
	RD = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	if err := RD.Ping(context.Background()).Err(); err != nil {
		log.Fatal(err)
	}
}
