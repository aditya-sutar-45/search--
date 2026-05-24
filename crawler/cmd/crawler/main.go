package main

import (
	"github.com/aditya-sutar-45/search--/crawler/crawl"
	"github.com/aditya-sutar-45/search--/crawler/frontier"
)

func main() {
	conn := frontier.NewRedisConnection("localhost:6379")
	seedURLs := []string{
		"https://en.wikipedia.org/wiki/Chicken",
		"https://detailed.com/50/",
	}

	c := crawl.NewCrawler(conn)
	c.Seed(seedURLs)

	c.StartCrawling()
}
