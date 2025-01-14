package seniverse

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"net/url"
	"sort"
	"strings"
)

func (c *V3Client) mapToUrl(urlStr string, params gorequest.Params) string {

	// 解析给定的 URL 字符串
	urlObj, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return ""
	}

	// 从 URL 中解析查询参数
	values, err := url.ParseQuery(urlObj.RawQuery)
	if err != nil {
		return ""
	}

	// 遍历 map，并将键值对添加到 url.Values 中
	for key, value := range params {
		values.Add(key, fmt.Sprint(value)) // 使用 fmt.Sprint 将值转换为字符串
	}

	// 将 url.Values 编码为查询字符串并更新 URL 对象的 RawQuery 字段
	urlObj.RawQuery = values.Encode()

	// 最终的 URL 字符串
	finalURL := urlObj.String()

	return finalURL
}

// 将 url.Value 转换成字符串
func (c *V4Client) urlValuesToString(v url.Values) string {
	if v == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]

		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(v)
		}
	}
	return buf.String()
}

func (c *V4Client) encodeUrlValues(v url.Values) string {
	if v == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]

		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(v))
		}
	}
	return buf.String()
}

// 签名
func (c *V4Client) sign(urlStr string, params gorequest.Params) string {

	// 解析给定的 URL 字符串
	urlObj, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return ""
	}

	// 从 URL 中解析查询参数
	values, err := url.ParseQuery(urlObj.RawQuery)
	if err != nil {
		return ""
	}

	// 遍历提供的参数并更新 URL 的查询参数
	for key, value := range params {
		values[key] = []string{fmt.Sprintf("%v", value)} // 使用字符串切片形式设置参数值
	}

	queryStr := c.urlValuesToString(values) // 将 url.Value 转换成字符串
	//encodeQueryStr := c.encodeUrlValues(values) // 编码Url值
	secretBytes := []byte(c.secret)

	hash := hmac.New(sha1.New, secretBytes)
	hash.Write([]byte(queryStr))

	encoded := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	encoded = url.QueryEscape(encoded)

	urlObj.RawQuery = fmt.Sprintf("%s&sig=%s", queryStr, encoded)
	//encodeUrl := fmt.Sprintf("%s://%s%s?%s&sig=%s", urlObj.Scheme, urlObj.Host, urlObj.Path, encodeQueryStr, encoded)

	return urlObj.String()
}
