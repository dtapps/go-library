package praise_goodness

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/dtapps/go-library/utils/gostring"
	"sort"
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
		if key == "notifyurl" {
			continue
		}
		signStr += fmt.Sprintf("%s=%s&", key, gostring.GetString(param.Get(key)))
	}
	signStr += "key=" + c.GetKey()

	return c.createSign(ctx, signStr)
}

// 签名
func (c *Client) createSign(ctx context.Context, signStr string) string {
	h := md5.New()
	h.Write([]byte(signStr))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
