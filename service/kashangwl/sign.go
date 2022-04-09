package kashangwl

import (
	"bytes"
	"crypto/md5"
	"dtapps/dta/library/utils/goparams"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

// md5(key + 参数1名称 + 参数1值 + 参数2名称 + 参数2值...) 加密源串应为{key}customer_id1192442order_id827669582783timestamp1626845767
func (app *App) getSign(customerKey string, params map[string]interface{}) string {
	// 参数按照参数名的字典升序排列
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	// 拼接
	query := bytes.NewBufferString(customerKey)
	for _, k := range keys {
		query.WriteString(k)
		query.WriteString(goparams.GetParamsString(params[k]))
	}
	// MD5加密
	h := md5.New()
	io.Copy(h, query)
	return strings.ToLower(hex.EncodeToString(h.Sum(nil)))
}

// 获取请求数据
func (app *App) getRequestData(params map[string]interface{}) string {
	// 公共参数
	args := url.Values{}
	// 请求参数
	for key, val := range params {
		args.Set(key, app.getString(val))
	}
	return args.Encode()
}

func (app *App) getString(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case bool:
		return strconv.FormatBool(v)
	default:
		marshal, _ := json.Marshal(v)
		return string(marshal)
	}
}
