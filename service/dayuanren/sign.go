package dayuanren

import (
	"fmt"
	"go.dtapp.net/library/utils/gomd5"
	"go.dtapp.net/library/utils/gorequest"
	"sort"
)

// 签名
// https://www.showdoc.com.cn/dyr/9227002900063946
func (c *Client) sign(param *gorequest.Params) string {
	var keys []string
	for k := range param.DeepGetAny() {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	signStr := ""
	for _, key := range keys {
		signStr += fmt.Sprintf("%s=%s&", key, gorequest.GetString(param.Get(key)))
	}
	signStr += fmt.Sprintf("apikey=%s", c.GetApiKey())
	return gomd5.ToUpper(signStr)
}
