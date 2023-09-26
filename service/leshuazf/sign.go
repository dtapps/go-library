package leshuazf

import (
	"github.com/dtapps/go-library/utils/gobase64"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gomd5"
	"github.com/dtapps/go-library/utils/gorequest"
	"sort"
)

// 数据签名
// https://www.yuque.com/leshuazf/doc/dbmxyi#Nr9Ps
func (c *Client) getSign(param gorequest.Params) string {
	// 参数按照参数名的字典升序排列
	var keys []string
	for k := range param {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	str, _ := gojson.Marshal(keys) // data字符串值
	return gobase64.Encode(gomd5.ToLower("lepos" + c.GetKeyAgent() + string(str)))
}
