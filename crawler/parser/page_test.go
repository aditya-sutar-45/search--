package parser

import (
	"fmt"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/aditya-sutar-45/search--/crawler/htmldownloader"
	"github.com/aditya-sutar-45/search--/crawler/utils"
)

func TestPageParsing(t *testing.T) {
	url := "https://en.wikipedia.org/wiki/Chicken"
	downloader := htmldownloader.NewHTMLDownloader()

	response, statusCode, contentType, err := downloader.Download(url)
	if err != nil {
		t.Fatal(err)
	}
	defer utils.CloseResponseBody(response)

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		t.Fatal(err)
	}

	parser := NewParser(document, url, statusCode, contentType)

	page := NewPage(
		parser.GetTitle(),
		parser.GetMetaDescription(),
		url,
		parser.GetNormalizedURLs(),
		parser.GetContent(),
		parser.GetDomain(),
		response.StatusCode,
		response.Header.Get("Content-Type"),
	)

	fmt.Println("===================================")
	fmt.Println("TITLE:", page.Title)
	fmt.Println("URL:", page.URL)
	fmt.Println("DOMAIN:", page.Domain)
	fmt.Println("STATUS CODE:", page.StatusCode)
	fmt.Println("CONTENT TYPE:", page.ContentType)
	fmt.Println("CRAWLED AT:", page.CrawledAt)
	fmt.Println("URL HASH:", page.URLHash)

	fmt.Println("\nCONTENT PREVIEW:")
	if len(page.Content) > 1000 {
		fmt.Println(page.Content[:1000] + "...")
	} else {
		fmt.Println(page.Content)
	}

	fmt.Println("\nLINK COUNT:", len(page.Links))

	fmt.Println("\nFIRST 10 LINKS:")
	for i, link := range page.Links {
		if i >= 10 {
			break
		}
		fmt.Printf("%d. %s\n", i+1, link)
	}

	fmt.Println("===================================")
}
