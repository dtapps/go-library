package jd

import (
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/dtapps/go-library/utils/gostring"
	"sort"
	"time"
)

func NewParamsWithType(_method string, params ...gorequest.Params) gorequest.Params {
	p := gorequest.NewParamsWith(params...)
	p.Set("method", _method)
	hh, _ := time.ParseDuration("8h")
	loc := time.Now().UTC().Add(hh)
	p.Set("timestamp", loc.Format("2006-01-02 15:04:05"))
	p.Set("format", "json")
	p.Set("v", "1.0")
	p.Set("sign_method", "md5")
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
	signStr := c.GetSecretKey()
	for _, key := range keys {
		signStr += key + gostring.GetString(p.Get(key))
	}
	signStr += c.GetSecretKey()
	p.Set("sign", createSign(signStr))
}

func SetCustomParameters(p gorequest.Params, uid string, sid string) gorequest.Params {
	p.Set("custom_parameters", map[string]interface{}{
		"uid": uid,
		"sid": sid,
	})
	return p
}
