// Package htmldownloader
package htmldownloader

import (
	"net/http"
	"time"
)

type HTMLDownloader struct {
	client  *http.Client
	headers map[string]string
}

func NewHTMLDownloader() *HTMLDownloader {
	client := &http.Client{Timeout: 10 * time.Second}

	headers := map[string]string{
		"User-Agent":      "SearchCrawler/1.0",
		"Accept-Language": "en-US,en;q=0.9",
	}

	return &HTMLDownloader{
		client:  client,
		headers: headers,
	}
}

func (h *HTMLDownloader) Download(url string) (*http.Response, int, string, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, 0, "", err
	}

	for key, value := range h.headers {
		request.Header.Set(key, value)
	}

	response, err := h.client.Do(request)
	if err != nil {
		return nil, 0, "", err
	}

	statusCode := response.StatusCode
	contentType := response.Header.Get("Content-Type")

	return response, statusCode, contentType, nil
}
