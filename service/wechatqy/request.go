package wechatqy

import "go.dtapp.net/library/utils/gorequest"

func (c *Client) request(url string, params map[string]interface{}) (gorequest.Response, error) {
	// 创建请求
	client := c.client

	// 设置请求地址
	client.SetUri(url)

	// 设置FORM格式
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
