package core

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func NewRedis(addr, password string, db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		panic(err)
	}

	return client
}
