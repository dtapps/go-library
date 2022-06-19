package eastiot

import (
	"go.dtapp.net/library/utils/gorequest"
	"time"
)

func (c *Client) request(url string, params map[string]interface{}, method string) (gorequest.Response, error) {

	// 公共参数
	params["appId"] = c.config.AppId
	params["timeStamp"] = time.Now().Unix()

	// 签名
	params["sign"] = c.getSign(params)

	// 创建请求
	client := c.client

	// 设置请求地址
	client.SetUri(url)

	// 设置方式
	client.SetMethod(method)

	// 设置格式
	client.SetContentTypeForm()

	// 设置参数
	client.SetParams(params)

	// 发起请求
	request, err := client.Request()
	if err != nil {
		return gorequest.Response{}, err
	}

	// 日志
	if c.logStatus == true {
		go c.postgresqlLog(request)
	}

	return request, err
}
