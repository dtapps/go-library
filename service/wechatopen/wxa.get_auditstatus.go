package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type GetAuditStatusResponse struct {
	Errcode    int    `json:"errcode"`    // 返回码
	Errmsg     string `json:"errmsg"`     // 错误信息
	Auditid    int    `json:"auditid"`    // 最新的审核 ID
	Status     int    `json:"status"`     // 审核状态
	Reason     string `json:"reason"`     // 当审核被拒绝时，返回的拒绝原因
	ScreenShot string `json:"ScreenShot"` // 当审核被拒绝时，会返回审核失败的小程序截图示例。用 | 分隔的 media_id 的列表，可通过获取永久素材接口拉取截图内容
}

type GetAuditStatusResult struct {
	Result GetAuditStatusResponse // 结果
	Body   []byte                 // 内容
	Http   gorequest.Response     // 请求
}

func newGetAuditStatusResult(result GetAuditStatusResponse, body []byte, http gorequest.Response) *GetAuditStatusResult {
	return &GetAuditStatusResult{Result: result, Body: body, Http: http}
}

// GetAuditStatus 查询审核单状态
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/code-management/getAuditStatus.html
func (c *Client) GetAuditStatus(ctx context.Context, authorizerAccessToken string, auditid int64, notMustParams ...gorequest.Params) (*GetAuditStatusResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "wxa/get_auditstatus")
	defer span.End()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("auditid", auditid)

	// 请求
	var response GetAuditStatusResponse
	request, err := c.request(ctx, span, "wxa/get_auditstatus?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newGetAuditStatusResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *GetAuditStatusResult) ErrcodeInfo() string {
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
