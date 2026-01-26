package config

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var RD *redis.Client

func initRedis() {
	RD = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if err := RD.Ping(context.Background()).Err(); err != nil {
		log.Fatal(err)
	}
}
