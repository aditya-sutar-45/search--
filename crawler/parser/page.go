package parser

import "time"

type Page struct {
	Title           string
	MetaDescription string
	URL             string
	Links           []string
	CrawledAt       time.Time
}

func NewPage(title string, metaDescription string, url string, links []string) *Page {
	return &Page{
		Title:           title,
		MetaDescription: metaDescription,
		URL:             url,
		Links:           links,
		CrawledAt:       time.Now(),
	}
}
