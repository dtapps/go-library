package wechatopen

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/gorequest"
	"net/http"
)

type WxaGetPageResponse struct {
	Errcode  int      `json:"errcode"`
	Errmsg   string   `json:"errmsg"`
	PageList []string `json:"page_list"` // page_list 页面配置列表
}

type WxaGetPageResult struct {
	Result WxaGetPageResponse // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
	Err    error              // 错误
}

func NewWxaGetPageResult(result WxaGetPageResponse, body []byte, http gorequest.Response, err error) *WxaGetPageResult {
	return &WxaGetPageResult{Result: result, Body: body, Http: http, Err: err}
}

// WxaGetPage 获取已上传的代码的页面列表
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_page.html
func (app *App) WxaGetPage() *WxaGetPageResult {
	// 请求
	request, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/wxa/get_page?access_token=%s", app.GetAuthorizerAccessToken()), map[string]interface{}{}, http.MethodGet)
	// 定义
	var response WxaGetPageResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewWxaGetPageResult(response, request.ResponseBody, request, err)
}
