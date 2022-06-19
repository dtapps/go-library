package pconline

import "go.dtapp.net/library/utils/gorequest"

func (c *Client) request(url string) (gorequest.Response, error) {

	// 创建请求
	client := c.client

	// 设置请求地址
	client.SetUri(url)

	// 发起请求
	request, err := client.Get()
	if err != nil {
		return gorequest.Response{}, err
	}

	// 日志
	if c.config.MongoDb != nil {
		go c.mongoLog(request)
	}
	if c.logStatus == true {
		go c.postgresqlLog(request)
	}

	return request, err
}
