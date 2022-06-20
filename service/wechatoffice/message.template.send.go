package wechatoffice

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
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
	Err    error                       // 错误
}

func newMessageTemplateSendResult(result MessageTemplateSendResponse, body []byte, http gorequest.Response, err error) *MessageTemplateSendResult {
	return &MessageTemplateSendResult{Result: result, Body: body, Http: http, Err: err}
}

// MessageTemplateSend 模板消息
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Template_Message_Interface.html
func (c *Client) MessageTemplateSend(notMustParams ...gorequest.Params) *MessageTemplateSendResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(fmt.Sprintf(apiUrl+"/cgi-bin/message/template/send?access_token=%s", c.getAccessToken()), params, http.MethodPost)
	// 定义
	var response MessageTemplateSendResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newMessageTemplateSendResult(response, request.ResponseBody, request, err)
}
