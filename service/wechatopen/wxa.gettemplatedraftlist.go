package wechatopen

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WxaGetTemplateDraftListResponse struct {
	Errcode   int    `json:"errcode"` // 返回码
	Errmsg    string `json:"errmsg"`  // 错误信息
	DraftList []struct {
		CreateTime             int64         `json:"create_time"`  // 开发者上传草稿时间戳
		UserVersion            string        `json:"user_version"` // 版本号，开发者自定义字段
		UserDesc               string        `json:"user_desc"`    // 版本描述   开发者自定义字段
		DraftId                int64         `json:"draft_id"`     // 草稿 id
		SourceMiniprogramAppid string        `json:"source_miniprogram_appid"`
		SourceMiniprogram      string        `json:"source_miniprogram"`
		Developer              string        `json:"developer"`
		CategoryList           []interface{} `json:"category_list"`
	} `json:"draft_list"` // 草稿信息列表
}

type WxaGetTemplateDraftListResult struct {
	Result WxaGetTemplateDraftListResponse // 结果
	Body   []byte                          // 内容
	Http   gorequest.Response              // 请求
}

func newWxaGetTemplateDraftListResult(result WxaGetTemplateDraftListResponse, body []byte, http gorequest.Response) *WxaGetTemplateDraftListResult {
	return &WxaGetTemplateDraftListResult{Result: result, Body: body, Http: http}
}

// WxaGetTemplateDraftList 获取代码草稿列表
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/code_template/gettemplatedraftlist.html
func (c *Client) WxaGetTemplateDraftList(ctx context.Context, notMustParams ...gorequest.Params) (*WxaGetTemplateDraftListResult, error) {
	// 检查
	err := c.checkComponentIsConfig()
	if err != nil {
		return nil, err
	}
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/wxa/gettemplatedraftlist?access_token=%s", c.GetComponentAccessToken(ctx)), params, http.MethodGet)
	if err != nil {
		return nil, err
	}
	// 定义
	var response WxaGetTemplateDraftListResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	if err != nil {
		return nil, err
	}
	return newWxaGetTemplateDraftListResult(response, request.ResponseBody, request), nil
}

// ErrcodeInfo 错误描述
func (resp *WxaGetTemplateDraftListResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 85064:
		return "找不到模板"
	}
	return "系统繁忙"
}
