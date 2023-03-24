package wechatpayopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type ProfitSharingReceiversDeleteResponse struct {
	SubMchid string `json:"sub_mchid"` // 子商户号
	Type     string `json:"type"`      // 分账接收方类型
	Account  string `json:"account"`   // 分账接收方账号
}

type ProfitSharingReceiversDeleteResult struct {
	Result   ProfitSharingReceiversDeleteResponse // 结果
	Body     []byte                               // 内容
	Http     gorequest.Response                   // 请求
	Err      error                                // 错误
	ApiError ApiError                             // 接口错误
}

func newProfitSharingReceiversDeleteResult(result ProfitSharingReceiversDeleteResponse, body []byte, http gorequest.Response, err error, apiError ApiError) *ProfitSharingReceiversDeleteResult {
	return &ProfitSharingReceiversDeleteResult{Result: result, Body: body, Http: http, Err: err, ApiError: apiError}
}

// ProfitSharingReceiversDelete 删除分账接收方API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter8_1_9.shtml
func (c *Client) ProfitSharingReceiversDelete(ctx context.Context, Type, account string) *ProfitSharingReceiversDeleteResult {
	// 参数
	params := gorequest.NewParams()
	params.Set("sub_mchid", c.GetSubMchId()) // 子商户号
	params.Set("appid", c.GetSpAppid())      // 应用ID
	params.Set("sub_appid", c.GetSubAppid()) // 子商户应用ID
	params.Set("type", Type)                 // 分账接收方类型
	if Type == MERCHANT_ID {
		params.Set("account", account) // 商户号
	}
	if Type == PERSONAL_OPENID {
		params.Set("account", account) // 个人openid
	}
	if Type == PERSONAL_SUB_OPENID {
		params.Set("account", account) // 个人sub_openid
	}
	// 请求
	request, err := c.request(ctx, apiUrl+"/v3/profitsharing/receivers/delete", params, http.MethodPost)
	if err != nil {
		return newProfitSharingReceiversDeleteResult(ProfitSharingReceiversDeleteResponse{}, request.ResponseBody, request, err, ApiError{})
	}
	// 定义
	var response ProfitSharingReceiversDeleteResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	// 错误
	var apiError ApiError
	err = gojson.Unmarshal(request.ResponseBody, &apiError)
	return newProfitSharingReceiversDeleteResult(response, request.ResponseBody, request, err, apiError)
}
