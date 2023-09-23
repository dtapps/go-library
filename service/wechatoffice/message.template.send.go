package wechatoffice

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type MessageTemplateSendResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Msgid   int    `json:"msgid"`
}

type MessageTemplateSendResult struct {
	Result MessageTemplateSendResponse // 结果
	Body   []byte                      // 内容
	Http   gorequest.Response          // 请求
}

func newMessageTemplateSendResult(result MessageTemplateSendResponse, body []byte, http gorequest.Response) *MessageTemplateSendResult {
	return &MessageTemplateSendResult{Result: result, Body: body, Http: http}
}

// MessageTemplateSend 模板消息
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html
func (c *Client) MessageTemplateSend(ctx context.Context, notMustParams ...*gorequest.Params) (*MessageTemplateSendResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/cgi-bin/message/template/send?access_token=%s", c.getAccessToken(ctx)), params, http.MethodPost)
	if err != nil {
		return newMessageTemplateSendResult(MessageTemplateSendResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response MessageTemplateSendResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newMessageTemplateSendResult(response, request.ResponseBody, request), err
}
