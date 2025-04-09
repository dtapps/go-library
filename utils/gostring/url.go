package gostring

import (
	"bytes"
	"strings"
)

const (
	prefixHTTP  = "http://"
	prefixHTTPS = "https://"
)

// CompleteUrlHttp 补全 URL
func CompleteUrlHttp(url string) string {
	if url == "" {
		return url
	}

	if strings.HasPrefix(url, prefixHTTP) {
		return url
	}

	var buffer bytes.Buffer

	if strings.HasPrefix(url, "//") {
		buffer.WriteString("http:")
		buffer.WriteString(url)
	} else if strings.HasPrefix(url, "://") {
		buffer.WriteString("http")
		buffer.WriteString(url)
	} else if strings.HasPrefix(url, prefixHTTPS) {
		buffer.WriteString(strings.Replace(url, prefixHTTPS, prefixHTTP, -1))
	} else {
		buffer.WriteString(prefixHTTP)
		buffer.WriteString(url)
	}

	return buffer.String()
}

// CompleteUrlHttps 补全 URL
func CompleteUrlHttps(url string) string {
	if url == "" {
		return url
	}

	if strings.HasPrefix(url, prefixHTTPS) {
		return url
	}

	var buffer bytes.Buffer

	if strings.HasPrefix(url, "//") {
		buffer.WriteString("https:")
		buffer.WriteString(url)
	} else if strings.HasPrefix(url, "://") {
		buffer.WriteString("https")
		buffer.WriteString(url)
	} else if strings.HasPrefix(url, prefixHTTP) {
		buffer.WriteString(strings.Replace(url, prefixHTTP, prefixHTTPS, -1))
	} else {
		buffer.WriteString(prefixHTTPS)
		buffer.WriteString(url)
	}

	return buffer.String()
}
