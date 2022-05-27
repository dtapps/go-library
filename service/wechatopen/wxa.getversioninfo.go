package wechatopen

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type WxaGetVersionInfoResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	ExpInfo struct {
		ExpTime    int64  `json:"exp_time"`    // 提交体验版的时间
		ExpVersion string `json:"exp_version"` // 体验版版本信息
		ExpDesc    string `json:"exp_desc"`    // 体验版版本描述
	} `json:"exp_info"` // 体验版信息
	ReleaseInfo struct {
		ReleaseTime    int64  `json:"release_time"`    // 发布线上版的时间
		ReleaseVersion string `json:"release_version"` // 线上版版本信息
		ReleaseDesc    string `json:"release_desc"`    // 线上版本描述
	} `json:"release_info"` // 线上版信息
}

type WxaGetVersionInfoResult struct {
	Result WxaGetVersionInfoResponse // 结果
	Body   []byte                    // 内容
	Http   gorequest.Response        // 请求
	Err    error                     // 错误
}

func NewWxaGetVersionInfoResult(result WxaGetVersionInfoResponse, body []byte, http gorequest.Response, err error) *WxaGetVersionInfoResult {
	return &WxaGetVersionInfoResult{Result: result, Body: body, Http: http, Err: err}
}

// WxaGetVersionInfo 查询小程序版本信息
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_versioninfo.html
func (app *App) WxaGetVersionInfo() *WxaGetVersionInfoResult {
	accessToken := app.GetAuthorizerAccessToken()
	if accessToken == "" {
		return NewWxaGetVersionInfoResult(WxaGetVersionInfoResponse{}, nil, gorequest.Response{}, errors.New("访问令牌为空"))
	}
	// 请求
	request, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/wxa/getversioninfo?access_token=%s", accessToken), map[string]interface{}{}, http.MethodPost)
	// 定义
	var response WxaGetVersionInfoResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewWxaGetVersionInfoResult(response, request.ResponseBody, request, err)
}
