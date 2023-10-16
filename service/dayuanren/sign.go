package dayuanren

import (
	"fmt"
	"github.com/dtapps/go-library/utils/gomd5"
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/dtapps/go-library/utils/gostring"
	"log"
	"sort"
)

// 签名
// https://www.showdoc.com.cn/dyr/9227002900063946
func (c *Client) sign(param gorequest.Params) string {
	var keys []string
	for k := range param {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	signStr := ""
	for _, key := range keys {
		signStr += fmt.Sprintf("%s=%s&", key, gostring.ToString(param.Get(key)))
	}
	log.Println(signStr)
	signStr += fmt.Sprintf("apikey=%s", c.GetApiKey())
	log.Println(signStr)
	return gomd5.ToUpper(signStr)
}
