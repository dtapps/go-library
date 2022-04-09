package wechatoffice

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DebugCgiBinTicketCheckResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type DebugCgiBinTicketCheckResult struct {
	Result DebugCgiBinTicketCheckResponse // 结果
	Body   []byte                         // 内容
	Err    error                          // 错误
}

func NewDebugCgiBinTicketCheckResult(result DebugCgiBinTicketCheckResponse, body []byte, err error) *DebugCgiBinTicketCheckResult {
	return &DebugCgiBinTicketCheckResult{Result: result, Body: body, Err: err}
}

// DebugCgiBinTicketCheck 判断ticket是否合法
// https://mp.weixin.qq.com/debug/cgi-bin/sandbox?t=jsapisign
func (app *App) DebugCgiBinTicketCheck() *DebugCgiBinTicketCheckResult {
	app.JsapiTicket = app.GetJsapiTicket()
	// 请求
	body, err := app.request(fmt.Sprintf("https://mp.weixin.qq.com/debug/cgi-bin/ticket/check?ticket=%s", app.JsapiTicket), map[string]interface{}{}, http.MethodGet)
	// 定义
	var response DebugCgiBinTicketCheckResponse
	err = json.Unmarshal(body, &response)
	return NewDebugCgiBinTicketCheckResult(response, body, err)
}
