package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type UndoAuditResponse struct {
	Errcode int    `json:"errcode"` // 返回码
	Errmsg  string `json:"errmsg"`  // 错误信息
}

type UndoAuditResult struct {
	Result UndoAuditResponse  // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newUndoAuditResult(result UndoAuditResponse, body []byte, http gorequest.Response) *UndoAuditResult {
	return &UndoAuditResult{Result: result, Body: body, Http: http}
}

// UndoAudit 撤回代码审核
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/code-management/undoAudit.html
func (c *Client) UndoAudit(ctx context.Context, authorizerAccessToken string, auditid int64, notMustParams ...gorequest.Params) (*UndoAuditResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("auditid", auditid)

	// 请求
	var response UndoAuditResponse
	request, err := c.request(ctx, "wxa/undocodeaudit?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newUndoAuditResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *UndoAuditResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 40001:
		return "获取 access_token 时 AppSecret 错误，或者 access_token 无效。请开发者认真比对 AppSecret 的正确性，或查看是否正在为恰当的公众号调用接口"
	case 87013:
		return "撤回次数达到上限（每天5次，每个月 10 次）"
	default:
		return resp.Result.Errmsg
	}
}
