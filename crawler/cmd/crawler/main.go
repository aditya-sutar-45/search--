package main

import (
	"log"

	"github.com/aditya-sutar-45/search--/crawler/crawl"
	"github.com/aditya-sutar-45/search--/crawler/frontier"
	"github.com/aditya-sutar-45/search--/crawler/storage"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoConn := storage.NewMongoConnection("search_engine")
	defer mongoConn.Disconnect()

	redisConn := frontier.NewRedisConnection("localhost:6379")

	pageRepo, err := storage.NewPageRepository(mongoConn)
	if err != nil {
		log.Fatalf("ERROR error creating page repository: %v\n", err)
	}

	seedURLs := []string{
		"https://en.wikipedia.org/wiki/Chicken",
		"https://detailed.com/50/",
	}

	c := crawl.NewCrawler(redisConn, pageRepo)
	c.Seed(seedURLs)

	c.StartCrawling()
}
