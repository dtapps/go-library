package pinduoduo

import (
	"encoding/json"
	"net/url"
	"sort"
	"strconv"
	"time"
)

// Params 请求参数
type Params map[string]interface{}

func NewParams() Params {
	p := make(Params)
	return p
}

func NewParamsWithType(_type string, params ...Params) Params {
	p := make(Params)
	p["type"] = _type
	p["timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	for _, v := range params {
		p.SetParams(v)
	}
	return p
}

func (app *App) Sign(p Params) {
	p["client_id"] = app.clientId
	// 排序所有的 key
	var keys []string
	for key := range p {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	signStr := app.clientSecret
	for _, key := range keys {
		signStr += key + getString(p[key])
	}
	signStr += app.clientSecret
	p["sign"] = createSign(signStr)
}

func (p Params) Set(key string, value interface{}) {
	p[key] = value
}

func (p Params) SetParams(params Params) {
	for key, value := range params {
		p[key] = value
	}
}

func (p Params) SetCustomParameters(uid string, sid string) {
	p["custom_parameters"] = map[string]interface{}{
		"uid": uid,
		"sid": sid,
	}
}

// SetGoodsSignList 设置商品goodsSign列表
func (p Params) SetGoodsSignList(goodsSign string) {
	p["goods_sign_list"] = []string{goodsSign}
}

func (p Params) GetQuery() string {
	u := url.Values{}
	for k, v := range p {
		u.Set(k, getString(v))
	}
	return u.Encode()
}

func getString(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case int:
		return strconv.Itoa(v)
	case bool:
		return strconv.FormatBool(v)
	default:

		bytes, _ := json.Marshal(v)
		return string(bytes)
	}
}
