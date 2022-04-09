package wechatopen

import (
	"encoding/json"
	"net/http"
)

type CgiBinComponentApiStartPushTicketResponse struct {
	AccessToken string `json:"access_token"` // 获取到的凭证
	ExpiresIn   int    `json:"expires_in"`   // 凭证有效时间，单位：秒。目前是7200秒之内的值
	Errcode     int    `json:"errcode"`      // 错误码
	Errmsg      string `json:"errmsg"`       // 错误信息
}

type CgiBinComponentApiStartPushTicketResult struct {
	Result CgiBinComponentApiStartPushTicketResponse // 结果
	Body   []byte                                    // 内容
	Err    error                                     // 错误
}

func NewCgiBinComponentApiStartPushTicketResult(result CgiBinComponentApiStartPushTicketResponse, body []byte, err error) *CgiBinComponentApiStartPushTicketResult {
	return &CgiBinComponentApiStartPushTicketResult{Result: result, Body: body, Err: err}
}

// CgiBinComponentApiStartPushTicket 启动ticket推送服务
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/component_verify_ticket_service.html
func (app *App) CgiBinComponentApiStartPushTicket() *CgiBinComponentApiStartPushTicketResult {
	// 参数
	param := NewParams()
	param["component_appid"] = app.ComponentAppId      // 平台型第三方平台的appid
	param["component_secret"] = app.ComponentAppSecret // 平台型第三方平台的APPSECRET
	params := app.NewParamsWith(param)
	// 请求
	body, err := app.request("https://api.weixin.qq.com/cgi-bin/component/api_start_push_ticket", params, http.MethodPost)
	// 定义
	var response CgiBinComponentApiStartPushTicketResponse
	err = json.Unmarshal(body, &response)
	return NewCgiBinComponentApiStartPushTicketResult(response, body, err)
}
