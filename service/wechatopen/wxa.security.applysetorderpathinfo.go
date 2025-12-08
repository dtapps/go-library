package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

// ApplySetOrderPathInfo 申请设置订单页path信息
// https://developers.weixin.qq.com/doc/oplatform/openApi/miniprogram-management/basic-info-management/api_applysetorderpathinfo.html
func (c *Client) ApplySetOrderPathInfo(ctx context.Context, batch_req map[string]any, notMustParams ...*gorequest.Params) (response APIResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("batch_req", batch_req) // 批量申请的信息

	// 请求
	err = c.request(ctx, "wxa/security/applysetorderpathinfo?access_token="+c.GetAuthorizerAccessToken(), params, http.MethodPost, &response)
	return
}

// ErrcodeInfo 错误描述
func GetApplySetOrderPathInfoErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 61042:
		return "批量提交超过最大数量，一次提交的 appid 数量不超过100个"
	case 61043:
		return "参数填写错误"
	case 61044:
		return "path填写不规范"
	default:
		return errmsg
	}
}
