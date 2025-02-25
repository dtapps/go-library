package meituan

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"
)

type SignHeaders struct {
	SCaApp              string
	SCaTimestamp        string
	SCaSignature        string
	SCaSignatureHeaders string
	ContentMD5          string
}

type SignUtil struct {
	AppKey    string
	AppSecret string
}

func NewSignUtil(appKey, appSecret string) *SignUtil {
	return &SignUtil{
		AppKey:    appKey,
		AppSecret: appSecret,
	}
}

func (su *SignUtil) GetSignHeaders(config map[string]interface{}) SignHeaders {
	signHeaders := SignHeaders{
		SCaApp:              su.AppKey,
		SCaTimestamp:        fmt.Sprintf("%d", time.Now().UnixNano()/int64(time.Millisecond)),
		SCaSignatureHeaders: "S-Ca-App,S-Ca-Timestamp",
		ContentMD5:          su.ContentMD5(config),
	}
	signHeaders.SCaSignature = su.Sign(config, signHeaders)
	return signHeaders
}

func (su *SignUtil) Sign(config map[string]interface{}, signHeaders SignHeaders) string {
	stringToSign := fmt.Sprintf("%s\n%s\n%s%s",
		su.HTTPMethod(config),
		signHeaders.ContentMD5,
		su.Headers(signHeaders),
		su.URL(config),
	)

	key := []byte(su.AppSecret)
	hash := su.hmacSHA256(key, []byte(stringToSign))
	return base64.StdEncoding.EncodeToString(hash)
}

func (su *SignUtil) hmacSHA256(key, data []byte) []byte {
	hmacSha256 := hmac.New(sha256.New, key)
	hmacSha256.Write(data)
	return hmacSha256.Sum(nil)
}

func (su *SignUtil) HTTPMethod(config map[string]interface{}) string {
	return strings.ToUpper(config["method"].(string))
}

func (su *SignUtil) ContentMD5(config map[string]interface{}) string {
	if config["method"] == "POST" && config["data"] != nil {
		bodyData, _ := json.Marshal(config["data"])
		hash := md5.Sum(bodyData)
		return base64.StdEncoding.EncodeToString(hash[:])
	}
	return ""
}

func (su *SignUtil) Headers(signHeaders SignHeaders) string {
	headers := map[string]string{
		"S-Ca-App":       signHeaders.SCaApp,
		"S-Ca-Timestamp": signHeaders.SCaTimestamp,
	}

	var keys []string
	for k := range headers {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var str string
	for _, k := range keys {
		str += fmt.Sprintf("%s:%s\n", k, headers[k])
	}
	return str
}

func (su *SignUtil) URL(config map[string]interface{}) string {

	praseURL, err := url.Parse(config["url"].(string))
	if err != nil {
		return ""
	}

	path := praseURL.Path
	if config["method"] == "get" {
		sortObj := su.objSort(config["data"].(map[string]interface{}))
		var keys []string
		for k := range sortObj {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		var str string
		for _, k := range keys {
			str += fmt.Sprintf("%s=%v&", k, sortObj[k])
		}
		path += "?" + str[:len(str)-1]

	}
	return path
}

func (su *SignUtil) objSort(arys map[string]interface{}) map[string]interface{} {
	var keys []string
	for k := range arys {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	newObj := make(map[string]interface{})
	for _, k := range keys {
		newObj[k] = arys[k]
	}
	return newObj
}
