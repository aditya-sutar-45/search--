// Package parser
package parser

import (
	"log"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Parser struct {
	Document *goquery.Document
	URL      string
}

func NewParser(doc *goquery.Document, url string) *Parser {
	return &Parser{
		Document: doc,
		URL:      url,
	}
}

func (p *Parser) GetTitle() string {
	return p.Document.Find("title").Text()
}

func (p *Parser) GetMetaDescription() string {
	desc, exits := p.Document.Find(`meta[name="description"]`).Attr("content")
	if !exits {
		return ""
	}

	return desc
}

func (p *Parser) GetNormalizedURLs() []string {
	urls := []string{}

	base, err := url.Parse(p.URL)
	if err != nil {
		log.Println("could not parse base urls")
		return urls
	}

	p.Document.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exits := s.Attr("href")
		if !exits {
			return
		}

		if href == "" ||
			strings.HasPrefix(href, "#") ||
			strings.HasPrefix(href, "javascript:") ||
			strings.HasPrefix(href, "mailto:") ||
			strings.HasPrefix(href, "tel:") {
			return
		}

		parsed, err := url.Parse(href)
		if err != nil {
			return
		}

		absolute := base.ResolveReference(parsed)
		normalizedURL, err := normalizeURL(absolute.String())
		if err != nil {
			return
		}

		urls = append(urls, normalizedURL)
	})

	return urls
}

func (p *Parser) Parse() *Page {
	pageTitle := p.GetTitle()
	metaDesc := p.GetMetaDescription()
	normalizedURLs := p.GetNormalizedURLs()

	page := NewPage(
		pageTitle,
		metaDesc,
		p.URL,
		normalizedURLs,
	)

	return page
}
