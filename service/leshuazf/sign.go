package leshuazf

import (
	"encoding/json"
	"go.dtapp.net/library/utils/gobase64"
	"go.dtapp.net/library/utils/gomd5"
	"sort"
)

// 数据签名
// https://www.yuque.com/leshuazf/doc/dbmxyi#Nr9Ps
func (app *App) getSign(params map[string]interface{}) string {
	// 参数按照参数名的字典升序排列
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	str, _ := json.Marshal(keys) // data字符串值
	return gobase64.Encode(gomd5.ToLower("lepos" + app.KeyAgent + string(str)))
}
