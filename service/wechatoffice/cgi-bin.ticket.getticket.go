package wechatoffice

import (
	"encoding/json"
	"fmt"
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
	Err    error                         // 错误
}

func NewCgiBinTicketGetTicketResult(result CgiBinTicketGetTicketResponse, body []byte, err error) *CgiBinTicketGetTicketResult {
	return &CgiBinTicketGetTicketResult{Result: result, Body: body, Err: err}
}

// CgiBinTicketGetTicket 获取api_ticket
// https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/JS-SDK.html
func (app *App) CgiBinTicketGetTicket(Type string) *CgiBinTicketGetTicketResult {
	app.AccessToken = app.GetAccessToken()
	// request
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=%s", app.AccessToken, Type), map[string]interface{}{}, http.MethodGet)
	// 定义
	var response CgiBinTicketGetTicketResponse
	err = json.Unmarshal(body, &response)
	return NewCgiBinTicketGetTicketResult(response, body, err)
}
