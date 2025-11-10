package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

// UndoAudit 撤回代码审核
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/code-management/undoAudit.html
func (c *Client) UndoAudit(ctx context.Context, authorizerAccessToken string, auditid int64, notMustParams ...*gorequest.Params) (response APIResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("auditid", auditid)

	// 请求
	err = c.request(ctx, "wxa/undocodeaudit?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return
}

// ErrcodeInfo 错误描述
func GetUndoAuditErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 40001:
		return "获取 access_token 时 AppSecret 错误，或者 access_token 无效。请开发者认真比对 AppSecret 的正确性，或查看是否正在为恰当的公众号调用接口"
	case 87013:
		return "撤回次数达到上限（每天5次，每个月 10 次）"
	default:
		return errmsg
	}
}
