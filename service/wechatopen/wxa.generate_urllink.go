package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type GenerateUrlLinkResponse struct {
	APIResponse        // 错误
	UrlLink     string `json:"url_link"` // 生成的小程序 URL Link
}

// GenerateUrlLink 获取加密URLLink
// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/qrcode-link/url-link/generateUrlLink.html
func (c *Client) GenerateUrlLink(ctx context.Context, authorizerAccessToken string, notMustParams ...*gorequest.Params) (response GenerateUrlLinkResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "wxa/generate_urllink?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return
}

// ErrcodeInfo 错误描述
func GetGenerateUrlLinkErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 40001:
		return "获取 access_token 时 AppSecret 错误，或者 access_token 无效。请开发者认真比对 AppSecret 的正确性，或查看是否正在为恰当的公众号调用接口"
	case 40002:
		return "暂无生成权限（个人主体小程序无权限，或者NFC 能力的小程序未申请权限）"
	case 40013:
		return "生成权限被封禁"
	case 85079:
		return "小程序没有线上版本，即小程序尚未发布，不可进行该操作"
	case 40165:
		return "参数path填写错误，更正后重试"
	case 40212:
		return "参数query填写错误 ，query格式遵循URL标准，即k1=v1&k2=v2"
	case 85401:
		return "参数expire_time填写错误，时间间隔大于1分钟且小于30天，更正后重试"
	case 85402:
		return "参数env_version填写错误，更正后重试"
	case 44990:
		return "频率过快，超过100次/秒；降低调用频率"
	case 44993:
		return "单天生成加密 URL Scheme+URL Link 数量超过上限50万"
	case 85406:
		return "URL Scheme（加密+明文）/加密 URL Link 单天累加访问次数超过上限"
	case 85088:
		return "请先开通云开发"
	default:
		return errmsg
	}
}
