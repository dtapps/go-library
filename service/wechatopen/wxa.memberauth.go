package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type GetTesterResponse struct {
	Errcode int    `json:"errcode"` // 错误码
	Errmsg  string `json:"errmsg"`  // 错误信息
	Members []struct {
		Userstr string `json:"userstr"` // 人员对应的唯一字符串
	} `json:"members"` // 人员信息列表
}

type GetTesterResult struct {
	Result GetTesterResponse  // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
}

func newGetTesterResult(result GetTesterResponse, body []byte, http gorequest.Response) *GetTesterResult {
	return &GetTesterResult{Result: result, Body: body, Http: http}
}

// GetTester 获取体验者列表
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/member-management/getTester.html
func (c *Client) GetTester(ctx context.Context, authorizerAccessToken string, notMustParams ...gorequest.Params) (*GetTesterResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "wxa/memberauth")
	defer span.End()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("action", "get_experiencer")

	// 请求
	var response GetTesterResponse
	request, err := c.request(ctx, span, "wxa/memberauth?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newGetTesterResult(response, request.ResponseBody, request), err
}
