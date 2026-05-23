package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"
)

func CloseResponseBody(r *http.Response) {
	if err := r.Body.Close(); err != nil {
		log.Println("error closing response body")
	}
}

func HashURL(url string) string {
	data := []byte(url)
	hash := sha256.Sum256(data)

	return hex.EncodeToString(hash[:])
}
