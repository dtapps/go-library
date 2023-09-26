package wechatminiprogram

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type SubscribeMessageSendResponse struct {
	Errcode int    // 错误码
	Errmsg  string // 错误信息
}

type SubscribeMessageSendResult struct {
	Result SubscribeMessageSendResponse // 结果
	Body   []byte                       // 内容
	Http   gorequest.Response           // 请求
}

func newSubscribeMessageSendResult(result SubscribeMessageSendResponse, body []byte, http gorequest.Response) *SubscribeMessageSendResult {
	return &SubscribeMessageSendResult{Result: result, Body: body, Http: http}
}

// SubscribeMessageSend 发送订阅消息
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/subscribe-message/subscribeMessage.send.html
func (c *Client) SubscribeMessageSend(ctx context.Context, notMustParams ...gorequest.Params) (*SubscribeMessageSendResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/cgi-bin/message/subscribe/send?access_token=%s", c.getAccessToken(ctx)), params, http.MethodPost)
	if err != nil {
		return newSubscribeMessageSendResult(SubscribeMessageSendResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response SubscribeMessageSendResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newSubscribeMessageSendResult(response, request.ResponseBody, request), err
}
