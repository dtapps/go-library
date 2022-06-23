package pinduoduo

import (
	"go.dtapp.net/library/utils/gorequest"
	"go.dtapp.net/library/utils/gostring"
)

func (c *Client) request(params map[string]interface{}) (gorequest.Response, error) {

	// 签名
	c.Sign(params)

	// 创建请求
	client := c.client

	// 设置参数
	client.SetParams(params)

	// 发起请求
	request, err := client.Get()
	if err != nil {
		return gorequest.Response{}, err
	}

	// 日志
	if c.config.PgsqlDb != nil {
		go c.log.GormMiddlewareCustom(gostring.ToString(params["type"]), request)
	}
	if c.config.MongoDb != nil {
		go c.log.MongoMiddlewareCustom(gostring.ToString(params["type"]), request)
	}

	return request, err
}
