// Package parser
package parser

import (
	"log"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Parser struct {
	Document    *goquery.Document
	RawHTML     string
	URL         string
	StatusCode  int
	ContentType string
}

func NewParser(doc *goquery.Document, rawHTML string, url string, statusCode int, contentType string) *Parser {
	return &Parser{
		Document:    doc,
		RawHTML:     rawHTML,
		URL:         url,
		StatusCode:  statusCode,
		ContentType: contentType,
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

func (p *Parser) GetContent() string {
	var content strings.Builder

	p.Document.Find(
		"script, style, noscript, nav, footer, header, aside",
	).Remove()

	p.Document.Find(
		"article p, main p, p, h1, h2, h3, h4, h5, h6",
	).Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())

		if text != "" {
			content.WriteString(text)
			content.WriteString(" ")
		}
	})

	return strings.Join(strings.Fields(content.String()), " ")
}

func (p *Parser) GetDomain() string {
	parsedURL, err := url.Parse(p.URL)
	if err != nil {
		return ""
	}

	return parsedURL.Host
}

func (p *Parser) Parse() *Page {
	pageTitle := p.GetTitle()
	metaDesc := p.GetMetaDescription()
	normalizedURLs := p.GetNormalizedURLs()
	domain := p.GetDomain()

	page := NewPage(
		pageTitle,
		metaDesc,
		p.URL,
		normalizedURLs,
		p.RawHTML,
		domain,
		p.StatusCode,
		p.ContentType,
	)

	return page
}
