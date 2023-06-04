package gostring

import (
	"strings"
)

// CompleteUrlHttp 补全 URL
func CompleteUrlHttp(url string) string {
	if url == "" {
		return url
	}
	if strings.HasPrefix(url, "//") {
		url = "http:" + url
	} else if strings.HasPrefix(url, "://") {
		url = "http" + url
	} else if strings.HasPrefix(url, "http://") {
	} else if strings.HasPrefix(url, "https://") {
		url = Replace(url, "https://", "http://")
	} else {
		url = "http://" + url
	}
	return url
}

// CompleteUrlHttps 补全 URL
func CompleteUrlHttps(url string) string {
	if url == "" {
		return url
	}
	if strings.HasPrefix(url, "//") {
		url = "https:" + url
	} else if strings.HasPrefix(url, "://") {
		url = "https" + url
	} else if strings.HasPrefix(url, "http://") {
		url = Replace(url, "http://", "https://")
	} else if strings.HasPrefix(url, "https://") {
	} else {
		url = "https://" + url
	}
	return url
}
