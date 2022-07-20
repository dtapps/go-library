package wechatopen

import (
	"encoding/json"
	"go.dtapp.net/library/utils/gorequest"
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
	Err    error                // 错误
}

func newTckWxPayListResult(result TckWxPayListResponse, body []byte, http gorequest.Response, err error) *TckWxPayListResult {
	return &TckWxPayListResult{Result: result, Body: body, Http: http, Err: err}
}

// TckWxPayList 获取授权绑定的商户号列表
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/cloudbase-common/wechatpay/getWechatPayList.html
func (c *Client) TckWxPayList() *TckWxPayListResult {
	// 参数
	params := gorequest.NewParams()
	// 请求
	request, err := c.request(apiUrl+"/tcb/wxpaylist?access_token="+c.GetComponentAccessToken(), params, http.MethodPost)
	// 定义
	var response TckWxPayListResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newTckWxPayListResult(response, request.ResponseBody, request, err)
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
