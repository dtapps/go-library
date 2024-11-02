package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type SetCodeAuditQuotaResponse struct {
	Errcode      int    `json:"errcode"`       // 错误码
	Errmsg       string `json:"errmsg"`        // 错误信息
	Rest         int64  `json:"rest"`          // quota剩余值
	Limit        int64  `json:"limit"`         // 当月分配quota
	SpeedupRest  int64  `json:"speedup_rest"`  // 剩余加急次数
	SpeedupLimit int64  `json:"speedup_limit"` // 当月分配加急次数
}

type SetCodeAuditQuotaResult struct {
	Result SetCodeAuditQuotaResponse // 结果
	Body   []byte                    // 内容
	Http   gorequest.Response        // 请求
}

func newSetCodeAuditQuotaResult(result SetCodeAuditQuotaResponse, body []byte, http gorequest.Response) *SetCodeAuditQuotaResult {
	return &SetCodeAuditQuotaResult{Result: result, Body: body, Http: http}
}

// SetCodeAuditQuota 查询服务商审核额度
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/code-management/setCodeAuditQuota.html
func (c *Client) SetCodeAuditQuota(ctx context.Context, authorizerAccessToken string, notMustParams ...gorequest.Params) (*SetCodeAuditQuotaResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response SetCodeAuditQuotaResponse
	request, err := c.request(ctx, "wxa/queryquota?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newSetCodeAuditQuotaResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *SetCodeAuditQuotaResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 40001:
		return "获取 access_token 时 AppSecret 错误，或者 access_token 无效。请开发者认真比对 AppSecret 的正确性，或查看是否正在为恰当的公众号调用接口"
	default:
		return resp.Result.Errmsg
	}
}
