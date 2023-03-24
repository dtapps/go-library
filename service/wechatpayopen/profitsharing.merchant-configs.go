package wechatpayopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type ProfitSharingMerchantConfigsResponse struct {
	SubMchid string `json:"sub_mchid"` // 子商户号
	MaxRatio int    `json:"max_ratio"` // 最大分账比例
}

type ProfitSharingMerchantConfigsResult struct {
	Result   ProfitSharingMerchantConfigsResponse // 结果
	Body     []byte                               // 内容
	Http     gorequest.Response                   // 请求
	Err      error                                // 错误
	ApiError ApiError                             // 接口错误
}

func newProfitSharingMerchantConfigsResult(result ProfitSharingMerchantConfigsResponse, body []byte, http gorequest.Response, err error, apiError ApiError) *ProfitSharingMerchantConfigsResult {
	return &ProfitSharingMerchantConfigsResult{Result: result, Body: body, Http: http, Err: err, ApiError: apiError}
}

// ProfitSharingMerchantConfigs 查询最大分账比例API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_1_7.shtml
func (c *Client) ProfitSharingMerchantConfigs(ctx context.Context) *ProfitSharingMerchantConfigsResult {
	// 参数
	params := gorequest.NewParams()
	// 请求
	request, err := c.request(ctx, apiUrl+"/v3/profitsharing/merchant-configs/"+c.GetSubMchId(), params, http.MethodGet)
	if err != nil {
		return newProfitSharingMerchantConfigsResult(ProfitSharingMerchantConfigsResponse{}, request.ResponseBody, request, err, ApiError{})
	}
	// 定义
	var response ProfitSharingMerchantConfigsResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	// 错误
	var apiError ApiError
	err = gojson.Unmarshal(request.ResponseBody, &apiError)
	return newProfitSharingMerchantConfigsResult(response, request.ResponseBody, request, err, apiError)
}
