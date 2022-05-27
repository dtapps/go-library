package gourl

import (
	"io"
	"net/http"
	"net/url"
	"os"
)

// LenCode 编码
func LenCode(s string) string {
	escape := url.QueryEscape(s)
	return escape
}

// DeCode 解码
func DeCode(s string) string {
	unescape, _ := url.QueryUnescape(s)
	return unescape
}

// ParseQuery 获取URL参数 https://studygolang.com/articles/2876
func ParseQuery(s string) map[string][]string {
	u, err := url.Parse(s)
	if err != nil {
		return nil
	}
	urlParam := u.RawQuery
	m, _ := url.ParseQuery(urlParam)
	return m
}

// QueryHeaders 获取Headers参数 https://blog.csdn.net/qq_31387691/article/details/109312920
func QueryHeaders(url string) map[string][]string {
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil
	}
	response, _ := client.Do(reqest)
	stdout := os.Stdout
	_, err = io.Copy(stdout, response.Body)
	headers := response.Header
	return headers
}
