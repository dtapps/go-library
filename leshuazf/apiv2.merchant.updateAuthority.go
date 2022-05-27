package leshuazf

import (
	"encoding/json"
	"go.dtapp.net/library/gorequest"
	"net/http"
)

type ApiV2MerchantUpdateAuthorityResponse struct {
}

type ApiV2MerchantUpdateAuthorityResult struct {
	Result ApiV2MerchantUpdateAuthorityResponse // 结果
	Body   []byte                               // 内容
	Http   gorequest.Response                   // 请求
	Err    error                                // 错误
}

func NewApiV2MerchantUpdateAuthorityResult(result ApiV2MerchantUpdateAuthorityResponse, body []byte, http gorequest.Response, err error) *ApiV2MerchantUpdateAuthorityResult {
	return &ApiV2MerchantUpdateAuthorityResult{Result: result, Body: body, Http: http, Err: err}
}

// ApiV2MerchantUpdateAuthority 给商户开通D0交易/结算权限接口。其中D0交易影响交易接口内t0字段能否标1，D0结算影响商户该种支付方式的秒到
// https://www.yuque.com/leshuazf/doc/dbmxyi#Vw97n
func (app *App) ApiV2MerchantUpdateAuthority(notMustParams ...Params) *ApiV2MerchantUpdateAuthorityResult {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	request, err := app.request("apiv2/merchant/updateAuthority", params, http.MethodPost)
	// 定义
	var response ApiV2MerchantUpdateAuthorityResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewApiV2MerchantUpdateAuthorityResult(response, request.ResponseBody, request, err)
}
