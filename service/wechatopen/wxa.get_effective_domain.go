package wechatopen

import (
	"context"
	"fmt"
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
	err := c.checkComponentIsConfig()
	if err != nil {
		return nil, err
	}
	err = c.checkAuthorizerIsConfig()
	if err != nil {
		return nil, err
	}
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/wxa/get_effective_domain?access_token=%s", c.GetAuthorizerAccessToken(ctx)), params, http.MethodPost)
	if err != nil {
		return nil, err
	}
	// 定义
	var response WxaGetEffectiveDomainResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		return nil, err
	}
	return newWxaGetEffectiveDomainResult(response, request.ResponseBody, request), nil
}
