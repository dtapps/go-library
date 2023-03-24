package wechatopen

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WxaApiWxAembeddedSetAuthorizeResponse struct {
	Errcode int    `json:"errcode"` // 返回码
	Errmsg  string `json:"errmsg"`  // 返回码信息
}

type WxaApiWxAembeddedSetAuthorizeResult struct {
	Result WxaApiWxAembeddedSetAuthorizeResponse // 结果
	Body   []byte                                // 内容
	Http   gorequest.Response                    // 请求
}

func newWxaApiWxAembeddedSetAuthorizeResult(result WxaApiWxAembeddedSetAuthorizeResponse, body []byte, http gorequest.Response) *WxaApiWxAembeddedSetAuthorizeResult {
	return &WxaApiWxAembeddedSetAuthorizeResult{Result: result, Body: body, Http: http}
}

// WxaApiWxAembeddedSetAuthorize 设置授权方式
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/embedded-management/setAuthorizedEmbedded.html
func (c *Client) WxaApiWxAembeddedSetAuthorize(ctx context.Context, notMustParams ...gorequest.Params) (*WxaApiWxAembeddedSetAuthorizeResult, error) {
	// 检查
	err := c.checkComponentIsConfig()
	if err != nil {
		return nil, err
	}
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/wxaapi/wxaembedded/set_authorize?access_token=%s", c.GetAuthorizerAccessToken(ctx)), params, http.MethodPost)
	if err != nil {
		return nil, err
	}
	// 定义
	var response WxaApiWxAembeddedSetAuthorizeResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		return nil, err
	}
	return newWxaApiWxAembeddedSetAuthorizeResult(response, request.ResponseBody, request), nil
}

// ErrcodeInfo 错误描述
func (resp *WxaApiWxAembeddedSetAuthorizeResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 89417:
		return "修改半屏小程序方式 flag 参数错误"
	}
	return "系统繁忙"
}
