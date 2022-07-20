package weishi

import "github.com/dtapps/go-library/utils/gorequest"

func (c *Client) request(url string) (gorequest.Response, error) {

	// 创建请求
	client := c.client

	// 设置请求地址
	client.SetUri(url)

	// 设置用户代理
	client.SetUserAgent(c.ua)

	// 发起请求
	request, err := client.Get()
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
