package movieapiv2

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strconv"
)

func (app *App) getSign(appSecret string, p map[string]interface{}) string {
	var keys []string
	for k := range p {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	signStr := ""
	for _, key := range keys {
		signStr += fmt.Sprintf("%s=%s&", key, app.getString(p[key]))
	}
	signStr += fmt.Sprintf("appSecret=%s", appSecret)
	// md5加密
	data := []byte(signStr)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
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
		bytes, _ := json.Marshal(v)
		return string(bytes)
	}
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
