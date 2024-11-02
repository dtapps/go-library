package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type UnbindTesterResponse struct {
	Errcode int    `json:"errcode"` // 错误码
	Errmsg  string `json:"errmsg"`  // 错误信息
}

type UnbindTesterResult struct {
	Result UnbindTesterResponse // 结果
	Body   []byte               // 内容
	Http   gorequest.Response   // 请求
}

func newUnbindTesterResult(result UnbindTesterResponse, body []byte, http gorequest.Response) *UnbindTesterResult {
	return &UnbindTesterResult{Result: result, Body: body, Http: http}
}

// UnbindTester 解除绑定体验者
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/member-management/unbindTester.html
func (c *Client) UnbindTester(ctx context.Context, authorizerAccessToken, wechatid, userstr string, notMustParams ...gorequest.Params) (*UnbindTesterResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	if wechatid != "" {
		params.Set("wechatid", wechatid)
	}
	params.Set("userstr", userstr)

	// 请求
	var response UnbindTesterResponse
	request, err := c.request(ctx, "wxa/unbind_tester?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newUnbindTesterResult(response, request.ResponseBody, request), err
}
