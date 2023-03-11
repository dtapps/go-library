package wechatopen

import (
	"context"
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type TckWxPayListResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	List    []struct {
		MerchantCode     string `json:"merchant_code"`
		MerchantName     string `json:"merchant_name"`
		CompanyName      string `json:"company_name"`
		MchRelationState string `json:"mch_relation_state"`
		JsapiAuthState   string `json:"jsapi_auth_state"`
		RefundAuthState  string `json:"refund_auth_state"`
	} `json:"list"`
}

type TckWxPayListResult struct {
	Result TckWxPayListResponse // 结果
	Body   []byte               // 内容
	Http   gorequest.Response   // 请求
}

func newTckWxPayListResult(result TckWxPayListResponse, body []byte, http gorequest.Response) *TckWxPayListResult {
	return &TckWxPayListResult{Result: result, Body: body, Http: http}
}

// TckWxPayList 获取授权绑定的商户号列表
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/cloudbase-common/wechatpay/getWechatPayList.html
func (c *Client) TckWxPayList(ctx context.Context, notMustParams ...gorequest.Params) (*TckWxPayListResult, error) {
	// 检查
	err := c.checkComponentIsConfig()
	if err != nil {
		return nil, err
	}
	// 参数
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/tcb/wxpaylist?access_token="+c.GetComponentAccessToken(ctx), params, http.MethodPost)
	if err != nil {
		return nil, err
	}
	// 定义
	var response TckWxPayListResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		return nil, err
	}
	return newTckWxPayListResult(response, request.ResponseBody, request), nil
}

// ErrcodeInfo 错误描述
func (resp *TckWxPayListResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 85064:
		return "找不到草稿"
	case 85065:
		return "模板库已满"
	}
	return "系统繁忙"
}
