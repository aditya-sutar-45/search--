package parser

import (
	"net/url"
	"path"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	if !strings.HasPrefix(rawURL, "http://") &&
		!strings.HasPrefix(rawURL, "https://") {
		rawURL = "https://" + rawURL
	}

	parsed, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	// Lowercase scheme and host
	parsed.Scheme = strings.ToLower(parsed.Scheme)
	parsed.Host = strings.ToLower(parsed.Host)

	// remove www.
	parsed.Host = strings.TrimPrefix(parsed.Host, "www.")

	// Remove fragments (#section)
	parsed.Fragment = ""

	// Remove default ports
	host := parsed.Hostname()
	port := parsed.Port()

	if (parsed.Scheme == "http" && port == "80") ||
		(parsed.Scheme == "https" && port == "443") {
		parsed.Host = host
	}

	// Normalize path
	parsed.Path = path.Clean(parsed.Path)

	if parsed.Path == "." {
		parsed.Path = "/"
	}

	// Remove trailing slash except root
	if parsed.Path != "/" {
		parsed.Path = strings.TrimRight(parsed.Path, "/")
	}

	// remove empty query params
	if parsed.RawQuery == "" {
		parsed.ForceQuery = false
	}

	// Remove tracking query params
	query := parsed.Query()

	trackingParams := []string{
		"utm_source",
		"utm_medium",
		"utm_campaign",
		"utm_term",
		"utm_content",
		"fbclid",
		"gclid",
	}

	for _, param := range trackingParams {
		query.Del(param)
	}

	// Sort query params automatically
	parsed.RawQuery = query.Encode()

	return parsed.String(), nil
}
