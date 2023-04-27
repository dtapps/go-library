package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WxaGetEffectiveDomainResponse struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	MpDomain struct {
		Requestdomain   []interface{} `json:"requestdomain"`
		Wsrequestdomain []interface{} `json:"wsrequestdomain"`
		Uploaddomain    []interface{} `json:"uploaddomain"`
		Downloaddomain  []interface{} `json:"downloaddomain"`
		Udpdomain       []interface{} `json:"udpdomain"`
		Tcpdomain       []interface{} `json:"tcpdomain"`
	} `json:"mp_domain"`
	ThirdDomain struct {
		Requestdomain   []interface{} `json:"requestdomain"`
		Wsrequestdomain []interface{} `json:"wsrequestdomain"`
		Uploaddomain    []interface{} `json:"uploaddomain"`
		Downloaddomain  []interface{} `json:"downloaddomain"`
		Udpdomain       []interface{} `json:"udpdomain"`
		Tcpdomain       []interface{} `json:"tcpdomain"`
	} `json:"third_domain"`
	DirectDomain struct {
		Requestdomain   []interface{} `json:"requestdomain"`
		Wsrequestdomain []interface{} `json:"wsrequestdomain"`
		Uploaddomain    []interface{} `json:"uploaddomain"`
		Downloaddomain  []interface{} `json:"downloaddomain"`
		Udpdomain       []interface{} `json:"udpdomain"`
		Tcpdomain       []interface{} `json:"tcpdomain"`
	} `json:"direct_domain"`
}

type WxaGetEffectiveDomainResult struct {
	Result WxaGetEffectiveDomainResponse // 结果
	Body   []byte                        // 内容
	Http   gorequest.Response            // 请求
}

func newWxaGetEffectiveDomainResult(result WxaGetEffectiveDomainResponse, body []byte, http gorequest.Response) *WxaGetEffectiveDomainResult {
	return &WxaGetEffectiveDomainResult{Result: result, Body: body, Http: http}
}

// WxaGetEffectiveDomain 获取发布后生效服务器域名列表
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_Basic_Info/get_effective_domain.html
func (c *Client) WxaGetEffectiveDomain(ctx context.Context, notMustParams ...gorequest.Params) (*WxaGetEffectiveDomainResult, error) {
	// 检查
	if err := c.checkAuthorizerConfig(ctx); err != nil {
		return newWxaGetEffectiveDomainResult(WxaGetEffectiveDomainResponse{}, []byte{}, gorequest.Response{}), err
	}
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/wxa/get_effective_domain?access_token="+c.GetAuthorizerAccessToken(ctx), params, http.MethodPost)
	if err != nil {
		return newWxaGetEffectiveDomainResult(WxaGetEffectiveDomainResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WxaGetEffectiveDomainResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWxaGetEffectiveDomainResult(response, request.ResponseBody, request), err
}
