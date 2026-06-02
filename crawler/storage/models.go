package storage

import (
	"time"

	"github.com/aditya-sutar-45/search--/crawler/parser"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type PageDocument struct {
	ID bson.ObjectID `bson:"_id,omitempty"`

	Title           string   `bson:"title"`
	MetaDescription string   `bson:"meta_description"`
	URL             string   `bson:"url"`
	URLHash         string   `bson:"url_hash"`
	Outlinks        []string `bson:"links"`

	RawHTML string `bson:"raw_html"`
	Domain  string `bson:"domain"`

	StatusCode  int       `bson:"status_code"`
	ContentType string    `bson:"content_type"`
	CrawledAt   time.Time `bson:"crawled_at"`
}

func pageToPageDocument(p *parser.Page) *PageDocument {
	return &PageDocument{
		Title:           p.Title,
		MetaDescription: p.MetaDescription,
		URL:             p.URL,
		URLHash:         p.URLHash,
		Outlinks:        p.Links,
		RawHTML:         p.RawHTMl,
		Domain:          p.Domain,
		StatusCode:      p.StatusCode,
		ContentType:     p.ContentType,
		CrawledAt:       p.CrawledAt,
	}
}
