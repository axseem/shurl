package utils

import (
	"net/url"
	"os"
	"strings"
)

const ALLOWED_CHARS = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"

func isOwnDomain(url string) bool {
	clearURL := strings.Replace(url, "http://", "", 1)
	clearURL = strings.Replace(clearURL, "https://", "", 1)
	clearURL = strings.Replace(clearURL, "www.", "", 1)
	clearURL = strings.Split(clearURL, "/")[0]

	return clearURL == os.Getenv("DOMAIN")
}

func IsValidURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return false
	}

	if isOwnDomain(u) {
		return false
	}

	return true
}

func IsValidID(id string) bool {
	for _, r := range id {
		if !strings.ContainsRune(ALLOWED_CHARS, r) {
			return false
		}
	}
	return true
}
