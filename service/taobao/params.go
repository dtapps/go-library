package taobao

import (
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/dtapps/go-library/utils/gostring"
	"sort"
	"strconv"
	"time"
)

func NewParamsWithType(_method string, params ...gorequest.Params) gorequest.Params {
	p := gorequest.NewParamsWith(params...)
	p.Set("method", _method)
	hh, _ := time.ParseDuration("8h")
	loc := time.Now().UTC().Add(hh)
	p.Set("timestamp", strconv.FormatInt(loc.Unix(), 10))
	p.Set("format", "json")
	p.Set("v", "2.0")
	p.Set("sign_method", "md5")
	//p.Set("partner_id", "Nilorg")
	for _, v := range params {
		p.SetParams(v)
	}
	return p
}

func (c *Client) Sign(p gorequest.Params) {
	p.Set("app_key", c.GetAppKey())
	// 排序所有的 key
	var keys []string
	for key := range p {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	signStr := c.GetAppSecret()
	for _, key := range keys {
		signStr += key + gostring.GetString(p.Get(key))
	}
	signStr += c.GetAppSecret()
	p.Set("sign", createSign(signStr))
}
