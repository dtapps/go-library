package pintoto

import (
	"go.dtapp.net/library/utils/gorequest"
	"time"
)

// 请求
func (c *Client) request(url string, params map[string]interface{}) (gorequest.Response, error) {

	// 公共参数
	params["time"] = time.Now().Unix()
	params["appKey"] = c.GetAppKey()

	// 签名
	params["sign"] = c.getSign(c.GetAppSecret(), params)

	// 创建请求
	client := c.client

	// 设置请求地址
	client.SetUri(url)

	// 设置格式
	client.SetContentTypeForm()

	// 设置参数
	client.SetParams(params)

	// 发起请求
	request, err := client.Post()
	if err != nil {
		return gorequest.Response{}, err
	}

	// 日志
	if c.config.PgsqlDb != nil {
		go c.log.GormMiddleware(request)
	}
	if c.config.MongoDb != nil {
		go c.log.MongoMiddleware(request)
	}

	return request, err
}
