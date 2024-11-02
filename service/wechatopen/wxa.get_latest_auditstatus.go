package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type GetLatestAuditStatusResponse struct {
	Errcode         int    `json:"errcode"`           // 返回码
	Errmsg          string `json:"errmsg"`            // 错误信息
	Auditid         int    `json:"auditid"`           // 最新的审核 ID
	Status          int    `json:"status"`            // 审核状态
	Reason          string `json:"reason"`            // 当审核被拒绝时，返回的拒绝原因
	ScreenShot      string `json:"ScreenShot"`        // 当审核被拒绝时，会返回审核失败的小程序截图示例。用 | 分隔的 media_id 的列表，可通过获取永久素材接口拉取截图内容
	UserDesc        string `json:"user_desc"`         // 审核版本
	UserVersion     string `json:"user_version"`      // 版本描述
	SubmitAuditTime int64  `json:"submit_audit_time"` // 时间戳，提交审核的时间
}

type GetLatestAuditStatusResult struct {
	Result GetLatestAuditStatusResponse // 结果
	Body   []byte                       // 内容
	Http   gorequest.Response           // 请求
}

func newGetLatestAuditStatusResult(result GetLatestAuditStatusResponse, body []byte, http gorequest.Response) *GetLatestAuditStatusResult {
	return &GetLatestAuditStatusResult{Result: result, Body: body, Http: http}
}

// GetLatestAuditStatus 查询最新一次审核单状态
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/code-management/getLatestAuditStatus.html
func (c *Client) GetLatestAuditStatus(ctx context.Context, authorizerAccessToken string, notMustParams ...gorequest.Params) (*GetLatestAuditStatusResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response GetLatestAuditStatusResponse
	request, err := c.request(ctx, "wxa/get_latest_auditstatus?access_token="+authorizerAccessToken, params, http.MethodGet, &response)
	return newGetLatestAuditStatusResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *GetLatestAuditStatusResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 86000:
		return "不是由第三方代小程序进行调用"
	case 86001:
		return "不存在第三方的已经提交的代码"
	case 85012:
		return "无效的审核 id"
	default:
		return resp.Result.Errmsg
	}
}
