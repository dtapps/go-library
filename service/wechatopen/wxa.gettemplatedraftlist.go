package wechatopen

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WxaGetTemplateDraftListResponse struct {
	Errcode   int    `json:"errcode"` // 返回码
	Errmsg    string `json:"errmsg"`  // 错误信息
	DraftList []struct {
		CreateTime             int           `json:"create_time"`  // 开发者上传草稿时间戳
		UserVersion            string        `json:"user_version"` // 版本号，开发者自定义字段
		UserDesc               string        `json:"user_desc"`    // 版本描述   开发者自定义字段
		DraftId                int           `json:"draft_id"`     // 草稿 id
		SourceMiniprogramAppid string        `json:"source_miniprogram_appid"`
		SourceMiniprogram      string        `json:"source_miniprogram"`
		Developer              string        `json:"developer"`
		CategoryList           []interface{} `json:"category_list"`
	} `json:"draft_list"` // 草稿信息列表
}

type WxaGetTemplateDraftListResult struct {
	Result WxaGetTemplateDraftListResponse // 结果
	Body   []byte                          // 内容
	Err    error                           // 错误
}

func NewWxaGetTemplateDraftListResult(result WxaGetTemplateDraftListResponse, body []byte, err error) *WxaGetTemplateDraftListResult {
	return &WxaGetTemplateDraftListResult{Result: result, Body: body, Err: err}
}

// WxaGetTemplateDraftList 获取代码草稿列表
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/code_template/gettemplatedraftlist.html
func (app *App) WxaGetTemplateDraftList() *WxaGetTemplateDraftListResult {
	app.componentAccessToken = app.GetComponentAccessToken()
	// 请求
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/wxa/gettemplatedraftlist?access_token=%s", app.componentAccessToken), map[string]interface{}{}, http.MethodGet)
	// 定义
	var response WxaGetTemplateDraftListResponse
	err = json.Unmarshal(body, &response)
	return NewWxaGetTemplateDraftListResult(response, body, err)
}
