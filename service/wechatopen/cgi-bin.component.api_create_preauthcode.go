package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type CgiBinComponentApiCreatePreAuthCodenResponse struct {
	PreAuthCode string `json:"pre_auth_code"` // 预授权码
	ExpiresIn   int64  `json:"expires_in"`    // 有效期，单位：秒
}

type CgiBinComponentApiCreatePreAuthCodenResult struct {
	Result CgiBinComponentApiCreatePreAuthCodenResponse // 结果
	Body   []byte                                       // 内容
	Http   gorequest.Response                           // 请求
}

func newCgiBinComponentApiCreatePreAuthCodenResult(result CgiBinComponentApiCreatePreAuthCodenResponse, body []byte, http gorequest.Response) *CgiBinComponentApiCreatePreAuthCodenResult {
	return &CgiBinComponentApiCreatePreAuthCodenResult{Result: result, Body: body, Http: http}
}

// CgiBinComponentApiCreatePreAuthCoden 预授权码
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/pre_auth_code.html
func (c *Client) CgiBinComponentApiCreatePreAuthCoden(ctx context.Context, componentAccessToken string, notMustParams ...*gorequest.Params) (*CgiBinComponentApiCreatePreAuthCodenResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("component_appid", c.config.componentAppId) // 第三方平台appid
	// 请求
	request, err := c.request(ctx, apiUrl+"/cgi-bin/component/api_create_preauthcode?component_access_token="+componentAccessToken, params, http.MethodPost)
	if err != nil {
		return newCgiBinComponentApiCreatePreAuthCodenResult(CgiBinComponentApiCreatePreAuthCodenResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response CgiBinComponentApiCreatePreAuthCodenResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newCgiBinComponentApiCreatePreAuthCodenResult(response, request.ResponseBody, request), err
}
