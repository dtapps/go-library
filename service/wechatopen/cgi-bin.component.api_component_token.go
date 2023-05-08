package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type CgiBinComponentApiComponentTokenResponse struct {
	ComponentAccessToken string `json:"component_access_token"` // 第三方平台 access_token
	ExpiresIn            int64  `json:"expires_in"`             // 有效期，单位：秒
}

type CgiBinComponentApiComponentTokenResult struct {
	Result CgiBinComponentApiComponentTokenResponse // 结果
	Body   []byte                                   // 内容
	Http   gorequest.Response                       // 请求
}

func newCgiBinComponentApiComponentTokenResult(result CgiBinComponentApiComponentTokenResponse, body []byte, http gorequest.Response) *CgiBinComponentApiComponentTokenResult {
	return &CgiBinComponentApiComponentTokenResult{Result: result, Body: body, Http: http}
}

// CgiBinComponentApiComponentToken 令牌
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/token/component_access_token.html
func (c *Client) CgiBinComponentApiComponentToken(ctx context.Context, notMustParams ...gorequest.Params) (*CgiBinComponentApiComponentTokenResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("component_appid", c.GetComponentAppId(ctx))                 // 第三方平台appid
	params.Set("component_appsecret", c.GetComponentAppSecret(ctx))         // 第三方平台appsecret
	params.Set("component_verify_ticket", GetComponentVerifyTicket(ctx, c)) // 微信后台推送的ticket
	// 请求
	request, err := c.request(ctx, apiUrl+"/cgi-bin/component/api_component_token", params, http.MethodPost)
	if err != nil {
		return newCgiBinComponentApiComponentTokenResult(CgiBinComponentApiComponentTokenResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response CgiBinComponentApiComponentTokenResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newCgiBinComponentApiComponentTokenResult(response, request.ResponseBody, request), err
}
