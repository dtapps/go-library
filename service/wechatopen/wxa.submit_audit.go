package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type SubmitAuditResponse struct {
	APIResponse       // 错误
	Auditid     int64 `json:"auditid"`
}

// SubmitAudit 提交代码审核
// https://developers.weixin.qq.com/doc/oplatform/openApi/miniprogram-management/code-management/api_submitaudit.html
func (c *Client) SubmitAudit(ctx context.Context, item_list map[string]any, notMustParams ...*gorequest.Params) (response SubmitAuditResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("item_list", item_list) // 审核项列表

	// 请求
	err = c.request(ctx, "wxa/submit_audit?access_token="+c.GetAuthorizerAccessToken(), params, http.MethodPost, &response)
	return
}
