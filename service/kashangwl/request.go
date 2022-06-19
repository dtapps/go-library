package kashangwl

import (
	"go.dtapp.net/library/utils/gorequest"
	"time"
)

func (c *Client) request(url string, params map[string]interface{}) (gorequest.Response, error) {

	// 公共参数
	params["timestamp"] = time.Now().UnixNano() / 1e6
	params["customer_id"] = c.GetCustomerId()

	// 签名参数
	params["sign"] = c.getSign(c.GetCustomerKey(), params)

	// 创建请求
	client := c.client

	// 设置请求地址
	client.SetUri(url)

	// 设置格式
	client.SetContentTypeJson()

	// 设置参数
	client.SetParams(params)

	// 发起请求
	request, err := client.Post()
	if err != nil {
		return gorequest.Response{}, err
	}

	// 日志
	if c.logStatus == true {
		go c.postgresqlLog(request)
	}

	return request, err
}
