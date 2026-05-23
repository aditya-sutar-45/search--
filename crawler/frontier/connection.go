// Package frontier
package frontier

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type ReddisConnection struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisConnection(redisAddr string) *ReddisConnection {
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	return &ReddisConnection{
		client: rdb,
		ctx:    context.Background(),
	}
}
