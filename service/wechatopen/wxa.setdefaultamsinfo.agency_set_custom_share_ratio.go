package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

// WxaSetDefaultamsInfoAgencySetCustomShareRatio
// 设置自定义分账比例
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/ams/percentage/SetCustomShareRatio.html
func (c *Client) WxaSetDefaultamsInfoAgencySetCustomShareRatio(ctx context.Context, appid string, shareRatio int64, notMustParams ...*gorequest.Params) (response APIRetResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("appid", appid)
	params.Set("share_ratio", shareRatio)

	// 请求
	err = c.request(ctx, "wxa/setdefaultamsinfo?action=agency_set_custom_share_ratio&access_token="+c.GetComponentAccessToken(), params, http.MethodPost, &response)
	return
}

// ErrcodeInfo 错误描述
func GetWxaSetDefaultamsInfoAgencySetCustomShareRatioErrcodeInfo(ret int, err_msg string) string {
	switch ret {
	case -202:
		return "内部错误"
	case 1700:
		return "参数错误"
	case 1701:
		return "参数错误"
	case 1737:
		return "操作过快"
	case 2056:
		return "服务商未在变现专区开通账户"
	default:
		return err_msg
	}
}
