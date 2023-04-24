package wnfuwu

import (
	"fmt"
	"github.com/dtapps/go-library/utils/gomd5"
	"github.com/dtapps/go-library/utils/gostring"
	"sort"
)

// 签名
func (c *Client) sign(params map[string]interface{}) string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	signStr := ""
	for _, key := range keys {
		signStr += fmt.Sprintf("%s=%s&", key, gostring.ToString(params[key]))
	}
	signStr += fmt.Sprintf("apikey=%s", c.GetApiKey())
	return gomd5.ToUpper(signStr)
}
