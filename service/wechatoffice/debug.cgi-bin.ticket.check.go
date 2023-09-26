package wechatoffice

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type DebugCgiBinTicketCheckResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type DebugCgiBinTicketCheckResult struct {
	Result DebugCgiBinTicketCheckResponse // 结果
	Body   []byte                         // 内容
	Http   gorequest.Response             // 请求
}

func newDebugCgiBinTicketCheckResult(result DebugCgiBinTicketCheckResponse, body []byte, http gorequest.Response) *DebugCgiBinTicketCheckResult {
	return &DebugCgiBinTicketCheckResult{Result: result, Body: body, Http: http}
}

// DebugCgiBinTicketCheck 判断ticket是否合法
// https://mp.weixin.qq.com/debug/cgi-bin/sandbox?t=jsapisign
func (c *Client) DebugCgiBinTicketCheck(ctx context.Context, notMustParams ...gorequest.Params) (*DebugCgiBinTicketCheckResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, fmt.Sprintf("https://mp.weixin.qq.com/debug/cgi-bin/ticket/check?ticket=%s", c.getJsapiTicket(ctx)), params, http.MethodGet)
	if err != nil {
		return newDebugCgiBinTicketCheckResult(DebugCgiBinTicketCheckResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response DebugCgiBinTicketCheckResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newDebugCgiBinTicketCheckResult(response, request.ResponseBody, request), err
}
