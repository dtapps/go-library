package meituan

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

type MediaClientNewSignResponse struct {
	SCaApp              string // 从网关申请的APP名称
	SCaSignature        string // 签名字符串（下面说明如何生成）
	SCaTimestamp        string // API 调用者传递时间戳，值为当前时间的毫秒数，也就是从1970年1月1日起至今的时间转换为毫秒，时间戳有效时间为2分钟(网关平台可以修改)。
	ContentMD5          string // Body MD5,服务端会校验Body内容是否被篡改 （下面有说明）
	SCaSignatureHeaders string // 将需要签名的header，使用英文逗号分割放到 Request 的 Header 中（下面有说明），其中必须包括S-Ca-Timestamp，建议将S-Ca-App也添加进去，例如【S-Ca-Timestamp,S-Ca-App】
}

// 调用方添加请求签名
// https://media.meituan.com/pc/index.html#/help?path=API%E6%8E%A5%E5%85%A5%E6%8C%87%E5%8D%97
func (c *MediaClient) NewSign(urlStr string, method string, param *gorequest.Params) MediaClientNewSignResponse {

	response := MediaClientNewSignResponse{
		SCaApp:              c.GetAppKey(),                                                    // 从网关申请的APP名称
		SCaTimestamp:        fmt.Sprintf("%d", time.Now().UnixNano()/int64(time.Millisecond)), // API 调用者传递时间戳，值为当前时间的毫秒数，也就是从1970年1月1日起至今的时间转换为毫秒，时间戳有效时间为2分钟(网关平台可以修改)。
		ContentMD5:          c.signContentMD5(method, param),
		SCaSignatureHeaders: "S-Ca-App,S-Ca-Timestamp",
	}
	response.SCaSignature = c.signSign(urlStr, method, param, response)

	return response
}

func (c *MediaClient) signSign(urlStr string, method string, param *gorequest.Params, response MediaClientNewSignResponse) string {
	stringToSign := fmt.Sprintf("%s\n%s\n%s%s",
		method,
		response.ContentMD5,
		c.signSort(method, param, response),
		c.signURL(urlStr, method, param),
	)

	key := []byte(c.GetAppSecret())
	hash := c.signHmacSHA256(key, []byte(stringToSign))
	return base64.StdEncoding.EncodeToString(hash)
}

func (c *MediaClient) signContentMD5(method string, param *gorequest.Params) string {
	if method == http.MethodPost && param != nil {
		bodyData, _ := json.Marshal(param.DeepGetAny())
		hash := md5.Sum(bodyData)
		return base64.StdEncoding.EncodeToString(hash[:])
	}
	return ""
}

func (c *MediaClient) signSort(method string, param *gorequest.Params, response MediaClientNewSignResponse) string {

	params := gorequest.NewParamsWith(param)
	params.Set("S-Ca-App", response.SCaApp)
	params.Set("S-Ca-Timestamp", response.SCaTimestamp)

	var keys []string
	for k := range params.DeepGetAny() {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var str string
	for _, k := range keys {
		str += fmt.Sprintf("%s:%s\n", k, params.Get(k))
	}
	return str
}

func (c *MediaClient) signURL(urlStr string, method string, param *gorequest.Params) string {

	praseURL, err := url.Parse(urlStr)
	if err != nil {
		return ""
	}

	path := praseURL.Path
	if method == http.MethodGet {
		sortParam := c.signURLSort(param)
		var keys []string
		for k := range sortParam.DeepGetAny() {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		var str string
		for _, k := range keys {
			str += fmt.Sprintf("%s=%v&", k, sortParam.Get(k))
		}
		path += "?" + str[:len(str)-1]

	}

	return path
}

func (c *MediaClient) signURLSort(param *gorequest.Params) *gorequest.Params {
	var keys []string
	for k := range param.DeepGetAny() {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	newParam := gorequest.NewParams()
	for _, k := range keys {
		newParam.Set(k, param.Get(k))
	}
	return newParam
}

func (c *MediaClient) signHmacSHA256(key, data []byte) []byte {
	hmacSha256 := hmac.New(sha256.New, key)
	hmacSha256.Write(data)
	return hmacSha256.Sum(nil)
}

// 计算 HMAC-SHA256 签名
func computeHMACSHA256(secret, data string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// 1.2 签名字符串说明
func prepareGenerateSignature(method, contentMD5, headersString, urlString string) string {
	stringToSign := ""
	if headersString == "" {
		stringToSign = fmt.Sprintf("%s\n%s\n%s%s", strings.ToUpper(method), contentMD5, headersString, urlString)
	} else {
		stringToSign = fmt.Sprintf("%s\n%s\n%s\n%s", strings.ToUpper(method), contentMD5, headersString, urlString)
	}
	return stringToSign
}

// 1.2 Content-MD5 说明
func computeMD5(param map[string]any) string {
	bodyData, _ := json.Marshal(param)
	hash := md5.Sum(bodyData)
	return base64.StdEncoding.EncodeToString(hash[:])
}

// 1.5 Headers 说明
func generateHeadersString(headers map[string]string) string {
	var keys []string
	for k := range headers {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var buffer bytes.Buffer
	for _, k := range keys {
		value := headers[k]
		if value == "" {
			buffer.WriteString(fmt.Sprintf("%s:\n", k))
		} else {
			buffer.WriteString(fmt.Sprintf("%s:%s\n", k, value))
		}
	}
	return buffer.String()
}

// 1.6 Url 说明
func generateURLString(path string, queryParams map[string]any) string {
	if len(queryParams) == 0 {
		return path
	}

	var keys []string
	for k := range queryParams {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var buffer bytes.Buffer
	buffer.WriteString(path)
	buffer.WriteString("?")

	for i, k := range keys {
		if queryParams[k] == "" {
			buffer.WriteString(k)
		} else {
			buffer.WriteString(fmt.Sprintf("%s=%s", k, url.QueryEscape(fmt.Sprintf("%v", queryParams[k]))))
		}
		if i < len(keys)-1 {
			buffer.WriteString("&")
		}
	}
	return buffer.String()
}

// 1.7 计算签名
func generateSignature(secret, stringToSign string) string {
	return computeHMACSHA256(secret, stringToSign)
}
