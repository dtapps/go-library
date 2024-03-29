package leshuazf

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type ApiV2MerchantUpdateAuthorityResponse struct {
}

type ApiV2MerchantUpdateAuthorityResult struct {
	Result ApiV2MerchantUpdateAuthorityResponse // 结果
	Body   []byte                               // 内容
	Http   gorequest.Response                   // 请求
}

func newApiV2MerchantUpdateAuthorityResult(result ApiV2MerchantUpdateAuthorityResponse, body []byte, http gorequest.Response) *ApiV2MerchantUpdateAuthorityResult {
	return &ApiV2MerchantUpdateAuthorityResult{Result: result, Body: body, Http: http}
}

// ApiV2MerchantUpdateAuthority 给商户开通D0交易/结算权限接口。其中D0交易影响交易接口内t0字段能否标1，D0结算影响商户该种支付方式的秒到
// https://www.yuque.com/leshuazf/doc/dbmxyi#Vw97n
func (c *Client) ApiV2MerchantUpdateAuthority(ctx context.Context, notMustParams ...gorequest.Params) (*ApiV2MerchantUpdateAuthorityResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, "/apiv2/merchant/updateAuthority", params, http.MethodPost)
	if err != nil {
		return newApiV2MerchantUpdateAuthorityResult(ApiV2MerchantUpdateAuthorityResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response ApiV2MerchantUpdateAuthorityResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newApiV2MerchantUpdateAuthorityResult(response, request.ResponseBody, request), err
}
