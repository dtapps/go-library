package wechatopen

import (
	"encoding/json"
	"fmt"
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
	Err    error              // 错误
}

func NewWxaGetPageResult(result WxaGetPageResponse, body []byte, err error) *WxaGetPageResult {
	return &WxaGetPageResult{Result: result, Body: body, Err: err}
}

// WxaGetPage 获取已上传的代码的页面列表
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/code/get_page.html
func (app *App) WxaGetPage() *WxaGetPageResult {
	app.authorizerAccessToken = app.GetAuthorizerAccessToken()
	// 请求
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/wxa/get_page?access_token=%s", app.authorizerAccessToken), map[string]interface{}{}, http.MethodGet)
	// 定义
	var response WxaGetPageResponse
	err = json.Unmarshal(body, &response)
	return NewWxaGetPageResult(response, body, err)
}
