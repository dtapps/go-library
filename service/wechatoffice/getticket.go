package wechatoffice

import (
	"encoding/json"
	"fmt"
)

// GetTicketResult 返回参数
type GetTicketResult struct {
	Errcode   int    `json:"errcode"` // 错误码
	Errmsg    string `json:"errmsg"`  // 错误信息
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
}

// GetTicket 获取api_ticket
// https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/JS-SDK.html
func (app *App) GetTicket(accessToken, Type string) (result GetTicketResult, err error) {
	// request
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=%s", accessToken, Type), map[string]interface{}{}, "GET")
	if err != nil {
		return
	}
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	return
}
