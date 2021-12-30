package wechatoffice

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GetTicketRespons struct {
	Errcode   int    `json:"errcode"` // 错误码
	Errmsg    string `json:"errmsg"`  // 错误信息
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
}

type GetTicketResult struct {
	Result GetTicketRespons // 结果
	Body   []byte           // 内容
	Err    error            // 错误
}

func NewGetTicketResult(result GetTicketRespons, body []byte, err error) *GetTicketResult {
	return &GetTicketResult{Result: result, Body: body, Err: err}
}

// GetTicket 获取api_ticket
// https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/JS-SDK.html
func (app *App) GetTicket(accessToken, Type string) *GetTicketResult {
	// request
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=%s", accessToken, Type), map[string]interface{}{}, http.MethodGet)
	// 定义
	var response GetTicketRespons
	err = json.Unmarshal(body, &response)
	return NewGetTicketResult(response, body, err)
}
