package wechatopen

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type WxaGetTemplateListResponse struct {
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
	TemplateList []struct {
		CreateTime             int           `json:"create_time"`              // 被添加为模板的时间
		UserVersion            string        `json:"user_version"`             // 模板版本号，开发者自定义字段
		UserDesc               string        `json:"user_desc"`                // 模板描述，开发者自定义字段
		TemplateId             int64         `json:"template_id"`              // 模板 id
		TemplateType           int           `json:"template_type"`            // 0对应普通模板，1对应标准模板
		SourceMiniprogramAppid string        `json:"source_miniprogram_appid"` // 开发小程序的appid
		SourceMiniprogram      string        `json:"source_miniprogram"`       // 开发小程序的名称
		Developer              string        `json:"developer"`                // 开发者
		CategoryList           []interface{} `json:"category_list"`
	} `json:"template_list"` // 模板信息列表
}

type WxaGetTemplateListResult struct {
	Result WxaGetTemplateListResponse // 结果
	Body   []byte                     // 内容
	Http   gorequest.Response         // 请求
	Err    error                      // 错误
}

func NewWxaGetTemplateListResult(result WxaGetTemplateListResponse, body []byte, http gorequest.Response, err error) *WxaGetTemplateListResult {
	return &WxaGetTemplateListResult{Result: result, Body: body, Http: http, Err: err}
}

// WxaGetTemplateList 获取代码模板列表
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/code_template/gettemplatelist.html
func (app *App) WxaGetTemplateList() *WxaGetTemplateListResult {
	// 请求
	request, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/wxa/gettemplatelist?access_token=%s", app.GetComponentAccessToken()), map[string]interface{}{}, http.MethodGet)
	// 定义
	var response WxaGetTemplateListResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewWxaGetTemplateListResult(response, request.ResponseBody, request, err)
}

// ErrcodeInfo 错误描述
func (resp *WxaGetTemplateListResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 43001:
		return "请使用GET，不要使用post"
	case 85064:
		return "找不到模板"
	}
	return "系统繁忙"
}
