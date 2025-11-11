package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type WxaGetTemplateDraftListResponse struct {
	APIResponse // 错误
	DraftList   []struct {
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

// WxaGetTemplateDraftList 获取代码草稿列表
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/code_template/gettemplatedraftlist.html
func (c *Client) WxaGetTemplateDraftList(ctx context.Context, notMustParams ...*gorequest.Params) (response WxaGetTemplateDraftListResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "wxa/gettemplatedraftlist?access_token="+c.GetComponentAccessToken(), params, http.MethodGet, &response)
	return
}

// ErrcodeInfo 错误描述
func GetWxaGetTemplateDraftListErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 85064:
		return "找不到模板"
	default:
		return errmsg
	}
}
