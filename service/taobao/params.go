package taobao

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

func NewParamsWithType(_method string, params ...Params) Params {
	p := make(Params)
	p["method"] = _method
	hh, _ := time.ParseDuration("8h")
	loc := time.Now().UTC().Add(hh)
	p["timestamp"] = strconv.FormatInt(loc.Unix(), 10)
	p["format"] = "json"
	p["v"] = "2.0"
	p["sign_method"] = "md5"
	//p["partner_id"] = "Nilorg"
	for _, v := range params {
		p.SetParams(v)
	}
	return p
}

func (c *Client) Sign(p Params) {
	p["app_key"] = c.GetAppKey()
	// 排序所有的 key
	var keys []string
	for key := range p {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	signStr := c.GetAppSecret()
	for _, key := range keys {
		signStr += key + getString(p[key])
	}
	signStr += c.GetAppSecret()
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
