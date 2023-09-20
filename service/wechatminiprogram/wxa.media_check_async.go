package wechatminiprogram

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WxaMediaCheckAsyncResponse struct {
	Errcode int    `json:"errcode"`  // 错误码
	Errmsg  string `json:"errmsg"`   // 错误信息
	TraceId string `json:"trace_id"` // 唯一请求标识，标记单次请求，用于匹配异步推送结果
	Result  struct {
		Suggest string `json:"suggest"` // 小程序的username
		Label   int    `json:"label"`   // 小程序的username
	} `json:"result"` // 小程序的username
	ToUserName   string `json:"ToUserName"`   // 小程序的username
	FromUserName string `json:"FromUserName"` // 平台推送服务UserName
	CreateTime   int    `json:"CreateTime"`   // 发送时间
	MsgType      string `json:"MsgType"`      // 默认为：event
	Event        string `json:"Event"`        // 默认为：wxa_media_check
	Appid        string `json:"appid"`        // 小程序的appid
	Version      int    `json:"version"`      // 小程序的username
	Detail       []struct {
		Strategy string `json:"strategy"` // 小程序的username
		Errcode  int    `json:"errcode"`  // 小程序的username
		Suggest  string `json:"suggest"`  // 小程序的username
		Label    int    `json:"label"`    // 小程序的username
		Prob     int    `json:"prob"`     // 小程序的username
	} `json:"detail"` // 小程序的username
}

type WxaMediaCheckAsyncResult struct {
	Result WxaMediaCheckAsyncResponse // 结果
	Body   []byte                     // 内容
	Http   gorequest.Response         // 请求
}

func newWxaMediaCheckAsyncResult(result WxaMediaCheckAsyncResponse, body []byte, http gorequest.Response) *WxaMediaCheckAsyncResult {
	return &WxaMediaCheckAsyncResult{Result: result, Body: body, Http: http}
}

// WxaMediaCheckAsync 音视频内容安全识别
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/sec-center/sec-check/mediaCheckAsync.html
func (c *Client) WxaMediaCheckAsync(ctx context.Context, notMustParams ...*gorequest.Params) (*WxaMediaCheckAsyncResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/wxa/media_check_async?access_token=%s", c.getAccessToken(ctx)), params, http.MethodPost)
	if err != nil {
		return newWxaMediaCheckAsyncResult(WxaMediaCheckAsyncResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WxaMediaCheckAsyncResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWxaMediaCheckAsyncResult(response, request.ResponseBody, request), err
}
