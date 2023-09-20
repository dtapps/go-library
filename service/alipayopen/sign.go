package alipayopen

import (
	"context"
	"crypto"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"log"
	"sort"
	"strconv"
	"strings"
)

func (c *Client) rsaSign(signContent string, hash crypto.Hash) (string, error) {

	shaNew := hash.New()
	shaNew.Write([]byte(signContent))
	hashed := shaNew.Sum(nil)

	signByte, err := c.privateKey.Sign(rand.Reader, hashed, crypto.SHA256)
	if err != nil {
		return "", err
	}

	sign := base64.StdEncoding.EncodeToString(signByte)

	return sign, nil
}

func (c *Client) sign(ctx context.Context, params *gorequest.Params) *gorequest.Params {
	// 排序
	var keys []string
	for key := range params.ToMap() {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// 转换成字符串
	var signStr = ""
	for _, key := range keys {
		if key != "" {
			signStr += fmt.Sprintf("%s=%s&", key, c.getString(params.Get(key)))
		}
	}

	signStr = strings.TrimRight(signStr, "&")

	// 签名
	sign, err := c.rsaSign(signStr, crypto.SHA256)
	if err != nil {
		log.Printf("签名失败：%s\n", err)
		return nil
	}
	params.Set("sign", sign)

	return params
}

func (c *Client) getString(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case int:
		return strconv.Itoa(v)
	case bool:
		return strconv.FormatBool(v)
	default:
		bytes, _ := gojson.Marshal(v)
		return string(bytes)
	}
}
