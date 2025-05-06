package x7s

import (
	"fmt"
	"go.dtapp.net/library/utils/gomd5"
	"go.dtapp.net/library/utils/gorequest"
	"log"
	"sort"
)

// 签名
func (c *Client) sign(param *gorequest.Params) string {
	var keys []string
	for k := range param.DeepGetAny() {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	signStr := ""
	for _, key := range keys {
		signStr += fmt.Sprintf("%s=%s&", key, gorequest.GetString(key))
	}
	signStr += fmt.Sprintf("key=%s", c.GetApiKey())
	log.Println(signStr)
	return gomd5.ToLower(signStr)
}
