package wechatpayopen

import (
	"context"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type ProfitSharingMerchantConfigsResponse struct {
	SubMchid string `json:"sub_mchid"` // 子商户号
	MaxRatio int    `json:"max_ratio"` // 最大分账比例
}

type ProfitSharingMerchantConfigsResult struct {
	Result ProfitSharingMerchantConfigsResponse // 结果
	Body   []byte                               // 内容
	Http   gorequest.Response                   // 请求
}

func newProfitSharingMerchantConfigsResult(result ProfitSharingMerchantConfigsResponse, body []byte, http gorequest.Response) *ProfitSharingMerchantConfigsResult {
	return &ProfitSharingMerchantConfigsResult{Result: result, Body: body, Http: http}
}

// ProfitSharingMerchantConfigs 查询最大分账比例API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_1_7.shtml
func (c *Client) ProfitSharingMerchantConfigs(ctx context.Context, notMustParams ...gorequest.Params) (*ProfitSharingMerchantConfigsResult, ApiError, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, fmt.Sprintf("v3/profitsharing/merchant-configs/%s", c.GetSubMchId()))
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response ProfitSharingMerchantConfigsResponse
	var apiError ApiError
	request, err := c.request(ctx, fmt.Sprintf("v3/profitsharing/merchant-configs/%s", c.GetSubMchId()), params, http.MethodGet, &response, &apiError)
	return newProfitSharingMerchantConfigsResult(response, request.ResponseBody, request), apiError, err
}
