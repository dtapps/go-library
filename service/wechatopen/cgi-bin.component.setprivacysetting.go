package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

// SetPrivacySetting 配置小程序用户隐私保护指引
// https://developers.weixin.qq.com/doc/oplatform/openApi/miniprogram-management/privacy-management/api_setprivacysetting.html
func (c *Client) SetPrivacySetting(ctx context.Context, notMustParams ...*gorequest.Params) (response APIResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "cgi-bin/component/setprivacysetting?access_token="+c.GetAuthorizerAccessToken(), params, http.MethodPost, &response)
	return
}

// GetSetPrivacySettingErrcodeInfo 错误描述
func GetSetPrivacySettingErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 86069:
		return "owner_setting必填字段字段缺失"
	case 86070:
		return "notice_method必填字段字段缺失"
	case 86072:
		return "store_expire_timestamp参数无效。如果是编码格式不对，也会报这个错"
	case 86073:
		return "ext_file_media_id参数无效"
	case 86074:
		return "现网隐私协议不存在"
	case 86075:
		return "现网隐私协议的ext_file_media_id禁止修改"
	default:
		return errmsg
	}
}
