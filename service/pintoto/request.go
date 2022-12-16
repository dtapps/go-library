package pintoto

import (
	"context"
	"github.com/dtapps/go-library"
	"github.com/dtapps/go-library/utils/gorequest"
	"time"
)

// 请求
func (c *Client) request(ctx context.Context, url string, params map[string]interface{}) (gorequest.Response, error) {

	// 公共参数
	params["time"] = time.Now().Unix()
	params["appKey"] = c.GetAppKey()

	// 签名
	params["sign"] = c.getSign(c.GetAppSecret(), params)

	// 创建请求
	client := c.requestClient

	// 设置请求地址
	client.SetUri(url)

	// 设置格式
	client.SetContentTypeForm()

	// 设置参数
	client.SetParams(params)

	// 发起请求
	request, err := client.Post(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 记录日志
	if c.log.status {
		go c.log.client.Middleware(ctx, request, go_library.Version())
	}

	return request, err
}
