package wechatopen

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type WxaGetAuditStatusResponse struct {
	Errcode    int    `json:"errcode"`    // 返回码
	Errmsg     string `json:"errmsg"`     // 错误信息
	Auditid    int    `json:"auditid"`    // 最新的审核 ID
	Status     int    `json:"status"`     // 审核状态
	Reason     string `json:"reason"`     // 当审核被拒绝时，返回的拒绝原因
	ScreenShot string `json:"ScreenShot"` // 当审核被拒绝时，会返回审核失败的小程序截图示例。用 | 分隔的 media_id 的列表，可通过获取永久素材接口拉取截图内容
}

type WxaGetAuditStatusResult struct {
	Result WxaGetAuditStatusResponse // 结果
	Body   []byte                    // 内容
	Http   gorequest.Response        // 请求
	Err    error                     // 错误
}

func newWxaGetAuditStatusResult(result WxaGetAuditStatusResponse, body []byte, http gorequest.Response, err error) *WxaGetAuditStatusResult {
	return &WxaGetAuditStatusResult{Result: result, Body: body, Http: http, Err: err}
}

// WxaGetAuditStatus 查询指定发布审核单的审核状态
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_auditstatus.html
func (c *Client) WxaGetAuditStatus(auditid int64) *WxaGetAuditStatusResult {
	// 参数
	params := gorequest.NewParams()
	params.Set("auditid", auditid)
	// 请求
	request, err := c.request(fmt.Sprintf(apiUrl+"/wxa/get_auditstatus?access_token=%s", c.GetAuthorizerAccessToken()), params, http.MethodPost)
	// 定义
	var response WxaGetAuditStatusResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newWxaGetAuditStatusResult(response, request.ResponseBody, request, err)
}

// ErrcodeInfo 错误描述
func (resp *WxaGetAuditStatusResult) ErrcodeInfo() string {
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
