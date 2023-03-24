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
func (c *Client) CgiBinComponentApiComponentToken(ctx context.Context) (*CgiBinComponentApiComponentTokenResult, error) {
	// 检查
	err := c.checkComponentIsConfig()
	if err != nil {
		return nil, err
	}
	// 参数
	param := gorequest.NewParams()
	param["component_appid"] = c.GetComponentAppId()                   // 第三方平台 appid
	param["component_appsecret"] = c.GetComponentAppSecret()           // 第三方平台 appsecret
	param["component_verify_ticket"] = c.GetComponentVerifyTicket(ctx) // 微信后台推送的 ticket
	params := gorequest.NewParamsWith(param)
	// 请求
	request, err := c.request(ctx, apiUrl+"/cgi-bin/component/api_component_token", params, http.MethodPost)
	if err != nil {
		return nil, err
	}
	// 定义
	var response CgiBinComponentApiComponentTokenResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		return nil, err
	}
	return newCgiBinComponentApiComponentTokenResult(response, request.ResponseBody, request), nil
}
