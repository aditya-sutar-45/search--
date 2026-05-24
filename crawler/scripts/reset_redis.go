package main

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

const (
	QUEUE_KEY       = "crawler_queue"
	VISITED_SET_KEY = "visited_urls"
)

func main() {
	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	err := client.Del(
		ctx,
		QUEUE_KEY,
		VISITED_SET_KEY,
	).Err()
	if err != nil {
		log.Fatalf("ERROR clearing redis keys: %v", err)
	}

	log.Println("SUCCESS cleared crawler queue and visited set")
}
