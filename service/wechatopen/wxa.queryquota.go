package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type SetCodeAuditQuotaResponse struct {
	APIResponse        // 错误
	Rest         int64 `json:"rest"`          // quota剩余值
	Limit        int64 `json:"limit"`         // 当月分配quota
	SpeedupRest  int64 `json:"speedup_rest"`  // 剩余加急次数
	SpeedupLimit int64 `json:"speedup_limit"` // 当月分配加急次数
}

// SetCodeAuditQuota 查询服务商审核额度
// https://developers.weixin.qq.com/doc/oplatform/openApi/miniprogram-management/code-management/api_setcodeauditquota.html
func (c *Client) SetCodeAuditQuota(ctx context.Context, notMustParams ...*gorequest.Params) (response SetCodeAuditQuotaResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "wxa/queryquota?access_token="+c.GetAuthorizerAccessToken(), params, http.MethodPost, &response)
	return
}

// ErrcodeInfo 错误描述
func GetSetCodeAuditQuotaErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 40001:
		return "获取 access_token 时 AppSecret 错误，或者 access_token 无效。请开发者认真比对 AppSecret 的正确性，或查看是否正在为恰当的公众号调用接口"
	default:
		return errmsg
	}
}
