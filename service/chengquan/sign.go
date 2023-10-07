package chengquan

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/dtapps/go-library/utils/gostring"
	"sort"
	"strings"
)

// 签名
func (c *Client) sign(ctx context.Context, param gorequest.Params) string {

	// 排序所有的 key
	var keys []string
	for key := range param {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	signStr := ""
	for _, key := range keys {
		if key != "" {
			signStr += fmt.Sprintf("%s=%s&", key, gostring.GetString(param.Get(key)))
		}
	}
	signStr += "key=" + c.GetAppKey()

	return c.createSign(signStr)
}

// 签名
func (c *Client) createSign(signStr string) string {
	h := md5.New()
	h.Write([]byte(signStr))
	cipherStr := h.Sum(nil)
	return strings.ToUpper(hex.EncodeToString(cipherStr))
}
