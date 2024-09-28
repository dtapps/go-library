package pinduoduo

import (
	"go.dtapp.net/library/utils/gorequest"
	"go.dtapp.net/library/utils/gostring"
	"sort"
	"strconv"
	"time"
)

func NewParamsWithType(_type string, params ...gorequest.Params) gorequest.Params {
	p := gorequest.NewParamsWith(params...)
	p.Set("type", _type)                                         // API接口名称
	p.Set("timestamp", strconv.FormatInt(time.Now().Unix(), 10)) // UNIX时间戳，单位秒，需要与拼多多服务器时间差值在10分钟内
	p.Set("data_type", "JSON")                                   // 响应格式，即返回数据的格式，JSON或者XML（二选一），默认JSON，注意是大写
	p.Set("version", "V1")                                       // 	API协议版本号，默认为V1，可不填
	for _, v := range params {
		p.SetParams(v)
	}
	return p
}

func (c *Client) Sign(p gorequest.Params) {
	p.Set("client_id", c.GetClientId()) // 	POP分配给应用的client_id
	// 排序所有的 key
	var keys []string
	for key := range p {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	signStr := c.GetClientSecret()
	for _, key := range keys {
		signStr += key + gostring.GetString(p.Get(key))
	}
	signStr += c.GetClientSecret()
	p.Set("sign", createSign(signStr))
}

func SetCustomParameters(p gorequest.Params, uid string, sid string) gorequest.Params {
	p.Set("custom_parameters", map[string]interface{}{
		"uid": uid,
		"sid": sid,
	})
	return p
}

// SetGoodsSignList 设置商品goodsSign列表
func SetGoodsSignList(p gorequest.Params, goodsSign string) gorequest.Params {
	p.Set("goods_sign_list", []string{goodsSign})
	return p
}
