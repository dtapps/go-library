package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

// ModifyJumpDomainDirectly 快速配置小程序业务域名
// https://developers.weixin.qq.com/doc/oplatform/openApi/miniprogram-management/domain-management/api_modifyjumpdomaindirectly.html
func (c *Client) ModifyJumpDomainDirectly(ctx context.Context, action string, webviewdomain []string, notMustParams ...*gorequest.Params) (response APIResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("action", action) // 操作类型
	if action != "get" {
		params.Set("webviewdomain", webviewdomain) // 小程序业务域名，当 action 参数是 get 时不需要此字段
	}

	// 请求
	err = c.request(ctx, "wxa/setwebviewdomain_directly?access_token="+c.GetAuthorizerAccessToken(), params, http.MethodPost, &response)
	return
}

// ErrcodeInfo 错误描述
func GetModifyJumpDomainDirectlyErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 86103:
		return "check confirmfile fail! 检查检验文件失败"
	case 506015:
		return "域名绑定的小程序超出上限"
	default:
		return errmsg
	}
}
