package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WxaMsgSecCheckResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Result  struct {
		Suggest string `json:"suggest"`
		Label   int    `json:"label"`
	} `json:"result"`
	Detail []struct {
		Strategy string `json:"strategy"`
		Errcode  int    `json:"errcode"`
		Suggest  string `json:"suggest"`
		Label    int    `json:"label"`
		Prob     int    `json:"prob,omitempty"`
		Level    int    `json:"level,omitempty"`
		Keyword  string `json:"keyword,omitempty"`
	} `json:"detail"`
	TraceId string `json:"trace_id"`
}

type WxaMsgSecCheckResult struct {
	Result WxaMsgSecCheckResponse // 结果
	Body   []byte                 // 内容
	Http   gorequest.Response     // 请求
}

func newWxaMsgSecCheckResult(result WxaMsgSecCheckResponse, body []byte, http gorequest.Response) *WxaMsgSecCheckResult {
	return &WxaMsgSecCheckResult{Result: result, Body: body, Http: http}
}

// WxaMsgSecCheck 文本内容安全识别
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html
func (c *Client) WxaMsgSecCheck(ctx context.Context, authorizerAccessToken string, notMustParams ...gorequest.Params) (*WxaMsgSecCheckResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/wxa/msg_sec_check?access_token="+authorizerAccessToken, params, http.MethodPost)
	if err != nil {
		return newWxaMsgSecCheckResult(WxaMsgSecCheckResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WxaMsgSecCheckResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWxaMsgSecCheckResult(response, request.ResponseBody, request), err
}
