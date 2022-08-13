package wechatoffice

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type CgiBinTicketGetTicketResponse struct {
	Errcode   int    `json:"errcode"` // 错误码
	Errmsg    string `json:"errmsg"`  // 错误信息
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
}

type CgiBinTicketGetTicketResult struct {
	Result CgiBinTicketGetTicketResponse // 结果
	Body   []byte                        // 内容
	Http   gorequest.Response            // 请求
	Err    error                         // 错误
}

func newCgiBinTicketGetTicketResult(result CgiBinTicketGetTicketResponse, body []byte, http gorequest.Response, err error) *CgiBinTicketGetTicketResult {
	return &CgiBinTicketGetTicketResult{Result: result, Body: body, Http: http, Err: err}
}

// CgiBinTicketGetTicket 获取api_ticket
// https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/JS-SDK.html
func (c *Client) CgiBinTicketGetTicket(ctx context.Context, Type string) *CgiBinTicketGetTicketResult {
	// request
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/cgi-bin/ticket/getticket?access_token=%s&type=%s", c.getAccessToken(ctx), Type), map[string]interface{}{}, http.MethodGet)
	// 定义
	var response CgiBinTicketGetTicketResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newCgiBinTicketGetTicketResult(response, request.ResponseBody, request, err)
}
