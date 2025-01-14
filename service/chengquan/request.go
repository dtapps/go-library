package chengquan

import (
	"context"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
	"go.dtapp.net/library/utils/gotime"
)

func (c *Client) request(ctx context.Context, url string, param *gorequest.Params, method string, response any) (gorequest.Response, error) {

	// 请求地址
	uri := c.GetApiURL() + url

	// 公共参数
	param.Set("timestamp", gotime.Current().TimestampWithMillisecond()) // 时间戳，以毫秒为单位。校验开发者与橙券的时间差，橙券允许开发者请求最大时间误差为3分钟 (3*60*1000)
	param.Set("app_id", c.GetAppID())                                   // 商户账号，由橙券提供，如：13105356515

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
	err = gojson.Unmarshal(request.ResponseBody, &response)

	return request, err
}
