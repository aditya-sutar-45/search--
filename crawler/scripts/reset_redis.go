package main

import (
	"context"
	"log"

	"github.com/aditya-sutar-45/search--/crawler/frontier"
	"github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	err := client.Del(
		ctx,
		frontier.QUEUE_KEY,
		frontier.VISITED_SET_KEY,
	).Err()
	if err != nil {
		log.Fatalf("ERROR clearing redis keys: %v", err)
	}

	log.Println("SUCCESS cleared crawler queue and visited set")
}
