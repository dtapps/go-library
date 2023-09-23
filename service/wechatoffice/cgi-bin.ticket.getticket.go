package wechatoffice

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
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
}

func newCgiBinTicketGetTicketResult(result CgiBinTicketGetTicketResponse, body []byte, http gorequest.Response) *CgiBinTicketGetTicketResult {
	return &CgiBinTicketGetTicketResult{Result: result, Body: body, Http: http}
}

// CgiBinTicketGetTicket 获取api_ticket
// https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/JS-SDK.html
func (c *Client) CgiBinTicketGetTicket(ctx context.Context, Type string, notMustParams ...*gorequest.Params) (*CgiBinTicketGetTicketResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/cgi-bin/ticket/getticket?access_token=%s&type=%s", c.getAccessToken(ctx), Type), params, http.MethodGet)
	if err != nil {
		return newCgiBinTicketGetTicketResult(CgiBinTicketGetTicketResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response CgiBinTicketGetTicketResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newCgiBinTicketGetTicketResult(response, request.ResponseBody, request), err
}
