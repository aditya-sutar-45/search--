package crawl

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/aditya-sutar-45/search--/crawler/frontier"
	"github.com/aditya-sutar-45/search--/crawler/htmldownloader"
	"github.com/aditya-sutar-45/search--/crawler/parser"
	"github.com/aditya-sutar-45/search--/crawler/utils"
)

type Crawler struct {
	queue       *frontier.Queue
	visistedSet *frontier.VisistedSet
	downloader  *htmldownloader.HTMLDownloader
}

func NewCrawler(r *frontier.ReddisConnection) *Crawler {
	return &Crawler{
		queue:       frontier.NewQueue(r),
		visistedSet: frontier.NewVisistedSet(r),
		downloader:  htmldownloader.NewHTMLDownloader(),
	}
}

func (c *Crawler) StartCrawling() {
	for {
		url, err := c.queue.Dequeue()
		if err != nil {
			log.Printf("ERROR dequeueing: %v\n", err)
			continue
		}

		err = c.Crawl(url)
		if err != nil {
			log.Printf("ERROR while crawling url %s: %v\n", url, err)
		}
	}
}

func (c *Crawler) Crawl(url string) error {
	response, statusCode, contentType, err := c.downloader.Download(url)
	if err != nil {
		return fmt.Errorf("error downloading content from url: %v", err)
	}
	defer utils.CloseResponseBody(response)

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return fmt.Errorf("error generating a goquery document: %v", err)
	}

	p := parser.NewParser(document, url, statusCode, contentType)
	page := p.Parse()

	log.Printf("INFO crawled: %s found %v URL's", url, len(page.Links))

	c.requeueURLs(page.Links)

	return nil
}

func (c *Crawler) requeueURLs(urls []string) {
	for _, url := range urls {
		if c.visistedSet.IsVisited(url) {
			continue
		}

		c.visistedSet.MarkVisisted(url)
		err := c.queue.Enqueue(url)
		if err != nil {
			log.Printf("ERROR could not requeue %s url: %v\n", url, err)
			continue
		}
	}
}

func (c *Crawler) Seed(seedURLs []string) {
	for _, url := range seedURLs {
		if c.visistedSet.IsVisited(url) {
			log.Printf("INFO already visisted %s\n", url)
			continue
		}
		err := c.queue.Enqueue(url)
		if err != nil {
			log.Printf("ERROR seeding url %s: %v\n", url, err)
			continue
		}
	}
}
