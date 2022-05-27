package wechatopen

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/gorequest"
	"net/http"
)

type WxaGetLatestAuditStatusResponse struct {
	Errcode    int    `json:"errcode"`    // 返回码
	Errmsg     string `json:"errmsg"`     // 错误信息
	Auditid    int    `json:"auditid"`    // 最新的审核 ID
	Status     int    `json:"status"`     // 审核状态
	Reason     string `json:"reason"`     // 当审核被拒绝时，返回的拒绝原因
	ScreenShot string `json:"ScreenShot"` // 当审核被拒绝时，会返回审核失败的小程序截图示例。用 | 分隔的 media_id 的列表，可通过获取永久素材接口拉取截图内容
}

type WxaGetLatestAuditStatusResult struct {
	Result WxaGetLatestAuditStatusResponse // 结果
	Body   []byte                          // 内容
	Http   gorequest.Response              // 请求
	Err    error                           // 错误
}

func NewWxaGetLatestAuditStatusResult(result WxaGetLatestAuditStatusResponse, body []byte, http gorequest.Response, err error) *WxaGetLatestAuditStatusResult {
	return &WxaGetLatestAuditStatusResult{Result: result, Body: body, Http: http, Err: err}
}

// WxaGetLatestAuditStatus 查询最新一次提交的审核状态
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_auditstatus.html
func (app *App) WxaGetLatestAuditStatus() *WxaGetLatestAuditStatusResult {
	// 请求
	request, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/wxa/get_latest_auditstatus?access_token=%s", app.GetAuthorizerAccessToken()), map[string]interface{}{}, http.MethodPost)
	// 定义
	var response WxaGetLatestAuditStatusResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewWxaGetLatestAuditStatusResult(response, request.ResponseBody, request, err)
}

// ErrcodeInfo 错误描述
func (resp *WxaGetLatestAuditStatusResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 86000:
		return "不是由第三方代小程序进行调用"
	case 86001:
		return "不存在第三方的已经提交的代码"
	case 85012:
		return "无效的审核 id"
	}
	return "系统繁忙"
}
