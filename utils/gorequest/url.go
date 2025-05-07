package gorequest

import (
	"net/url"
	"strings"
)

type UriParse struct {
	uri string
}

func NewUri(uri string) *UriParse {
	return &UriParse{uri: uri}
}

// ResponseUrlParse 返回参数
type ResponseUrlParse struct {
	Uri      string `json:"uri"`       // URI
	Urn      string `json:"urn"`       // URN
	Url      string `json:"url"`       // URL
	Scheme   string `json:"scheme"`    // 协议
	Host     string `json:"host"`      // 主机
	Hostname string `json:"hostname"`  // 主机名
	Port     string `json:"port"`      // 端口
	Path     string `json:"path"`      // 路径
	RawQuery string `json:"raw_query"` // 参数 ?
	Fragment string `json:"fragment"`  // 片段 #
}

// Parse 解析URl
func (u *UriParse) Parse() (resp ResponseUrlParse) {
	parse, err := url.Parse(u.uri)
	if err != nil {
		return
	}
	resp.Uri = u.uri
	resp.Urn = parse.Host + parse.Path
	resp.Url = parse.Scheme + "://" + parse.Host + parse.Path
	resp.Scheme = parse.Scheme
	resp.Host = parse.Host
	resp.Hostname = parse.Hostname()
	resp.Port = parse.Port()
	resp.Path = parse.Path
	resp.RawQuery = parse.RawQuery
	resp.Fragment = parse.Fragment
	return
}

// UriFilterExcludeQueryString 过滤掉url中的参数
func (u *UriParse) UriFilterExcludeQueryString() string {
	URL, _ := url.Parse(u.uri)
	clearUri := strings.ReplaceAll(u.uri, URL.RawQuery, "")
	clearUri = strings.TrimRight(clearUri, "?")
	return strings.TrimRight(clearUri, "/")
}

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

func IsHttpURL(s string) bool {
	_, err := url.ParseRequestURI(s)
	if err != nil {
		// 解析失败，不是有效的URL
		return false
	}
	return true
}

func IsURL(s string) bool {
	u, err := url.Parse(s)
	if err != nil {
		// 解析失败，不是有效的URL
		return false
	}
	// 检查是否有方案（scheme），例如 http, https 等
	return u.Scheme != ""
}

// ParseQueryString 解析 URL 查询字符串为 map 类型。
func ParseQueryString(input string) map[string]any {
	paramMap := make(map[string]interface{})
	keyValuePairs := strings.Split(input, "&")
	for _, pair := range keyValuePairs {
		parts := strings.Split(pair, "=")
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			paramMap[key] = value
		}
	}
	return paramMap
}
