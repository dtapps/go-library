package praise_goodness

import (
	"context"
	"encoding/json"
	"go.dtapp.net/library/utils/gorequest"
	"strconv"
	"time"
)

func (c *Client) request(ctx context.Context, url string, param *gorequest.Params, method string, response any) (gorequest.Response, error) {

	// 请求地址
	uri := c.GetApiURL() + url

	// 公共参数
	param.Set("times", strconv.FormatInt(time.Now().Unix(), 10)) // 创建时间，秒级时间戳

	// 签名
	param.Set("sign", c.sign(ctx, param))

	// 设置请求地址
	c.httpClient.SetUri(uri)

	// 设置方式
	c.httpClient.SetMethod(method)

	// 设置格式
	c.httpClient.SetContentTypeForm()

	// 设置参数
	c.httpClient.SetParams(param)

	// 发起请求
	request, err := c.httpClient.Request(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 解析响应
	err = json.Unmarshal(request.ResponseBody, &response)

	return request, err
}
