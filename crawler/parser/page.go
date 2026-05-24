package parser

import (
	"time"

	"github.com/aditya-sutar-45/search--/crawler/utils"
)

type Page struct {
	Title           string
	MetaDescription string
	URL             string
	URLHash         string
	Links           []string

	Content string
	Domain  string

	StatusCode  int
	ContentType string

	CrawledAt time.Time
}

func NewPage(
	title string,
	metaDescription string,
	url string,
	links []string,
	content string,
	domain string,
	statusCode int,
	contentType string,
) *Page {
	urlHash := utils.HashURL(url)

	return &Page{
		Title:           title,
		MetaDescription: metaDescription,
		URL:             url,
		URLHash:         urlHash,
		Links:           links,

		Content:     content,
		Domain:      domain,
		StatusCode:  statusCode,
		ContentType: contentType,

		CrawledAt: time.Now(),
	}
}
