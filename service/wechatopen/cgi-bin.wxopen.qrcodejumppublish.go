package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

// QrcodeJumpPublish 发布已设置的二维码规则
// https://developers.weixin.qq.com/doc/oplatform/openApi/miniprogram-management/jumpqrcode-config/api_qrcodejumppublish.html
func (c *Client) QrcodeJumpPublish(ctx context.Context, prefix string, notMustParams ...*gorequest.Params) (response APIResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("prefix", prefix)

	// 请求
	err = c.request(ctx, "cgi-bin/wxopen/qrcodejumppublish?access_token="+c.GetAuthorizerAccessToken(), params, http.MethodPost, &response)
	return
}

// ErrcodeInfo 错误描述
func GetQrcodeJumpPublishErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 44990:
		return "接口请求太快（超过5次/秒）"
	case 85074:
		return "小程序未发布, 小程序必须先发布代码才可以发布二维码跳转规则"
	case 85075:
		return "个人类型小程序无法设置二维码规则"
	case 85095:
		return "数据异常，请删除后重新添加"
	case 886000:
		return "本月发布次数达到上线（100次）"
	default:
		return errmsg
	}
}
