package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
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
}

func newWxaGetTemplateListResult(result WxaGetTemplateListResponse, body []byte, http gorequest.Response) *WxaGetTemplateListResult {
	return &WxaGetTemplateListResult{Result: result, Body: body, Http: http}
}

// WxaGetTemplateList 获取代码模板列表
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/code_template/gettemplatelist.html
func (c *Client) WxaGetTemplateList(ctx context.Context, componentAccessToken string, notMustParams ...*gorequest.Params) (*WxaGetTemplateListResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/wxa/gettemplatelist?access_token="+componentAccessToken, params, http.MethodGet)
	if err != nil {
		return newWxaGetTemplateListResult(WxaGetTemplateListResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WxaGetTemplateListResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWxaGetTemplateListResult(response, request.ResponseBody, request), err
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
