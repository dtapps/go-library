package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type WxaMsgSecCheckResponse struct {
	APIResponse // 错误
	Result      struct {
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

// WxaMsgSecCheck 文本内容安全识别
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.getUnlimited.html
func (c *Client) WxaMsgSecCheck(ctx context.Context, authorizerAccessToken string, notMustParams ...*gorequest.Params) (response WxaMsgSecCheckResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "wxa/msg_sec_check?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return
}
