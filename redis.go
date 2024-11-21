package core

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var G_REDIS *redis.Client

func NewRedis(addr, password string, db int) {
	G_REDIS = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	if _, err := G_REDIS.Ping(context.Background()).Result(); err != nil {
		panic(err)
	}
}
