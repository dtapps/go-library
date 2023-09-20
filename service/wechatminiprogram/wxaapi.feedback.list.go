package wechatminiprogram

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WxaApiFeedbackListResponse struct {
	List []struct {
		RecordId   int      `json:"record_id"`
		CreateTime int      `json:"create_time"`
		Content    string   `json:"content"`
		Phone      string   `json:"phone"`
		Openid     string   `json:"openid"`
		Nickname   string   `json:"nickname"`
		HeadUrl    string   `json:"head_url"`
		Type       int      `json:"type"`
		MediaIds   []string `json:"mediaIds"`
	} `json:"list"`
	TotalNum int `json:"total_num"`
	Errcode  int `json:"errcode"`
	RpcCount int `json:"__rpcCount"`
}

type WxaApiFeedbackListResult struct {
	Result WxaApiFeedbackListResponse // 结果
	Body   []byte                     // 内容
	Http   gorequest.Response         // 请求
	Err    error                      // 错误
}

func newWxaApiFeedbackListResult(result WxaApiFeedbackListResponse, body []byte, http gorequest.Response, err error) *WxaApiFeedbackListResult {
	return &WxaApiFeedbackListResult{Result: result, Body: body, Http: http, Err: err}
}

// WxaApiFeedbackList 获取用户反馈列表
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/operation/operation.getFeedback.html
func (c *Client) WxaApiFeedbackList(ctx context.Context, notMustParams ...*gorequest.Params) *WxaApiFeedbackListResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("access_token", c.getAccessToken(ctx))
	// 请求
	request, err := c.request(ctx, apiUrl+"/wxaapi/feedback/list", params, http.MethodGet)
	// 定义
	var response WxaApiFeedbackListResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWxaApiFeedbackListResult(response, request.ResponseBody, request, err)
}
