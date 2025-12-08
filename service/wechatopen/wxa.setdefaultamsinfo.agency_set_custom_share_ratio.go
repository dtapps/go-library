package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

// SetCustomShareRatio 设置自定义分账比例
// https://developers.weixin.qq.com/doc/oplatform/openApi/ams/percentage/api_setcustomshareratio.html
func (c *Client) SetCustomShareRatio(ctx context.Context, appid string, shareRatio int64, notMustParams ...*gorequest.Params) (response APIRetResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("share_ratio", shareRatio) // 服务商自定义分账比例。签约时，默认优先使用自定义分账比例，若不存在，则使用默认分账比例。如share_ratio为40，则代表服务商获得收益的40%，小程序商家获得收益的60%
	params.Set("appid", appid)            // 针对该小程序APPID设置自定义分账比例。

	// 请求
	err = c.request(ctx, "wxa/setdefaultamsinfo?action=agency_set_custom_share_ratio&access_token="+c.GetComponentAccessToken(), params, http.MethodPost, &response)
	return
}

// ErrcodeInfo 错误描述
func GetSetCustomShareRatioErrcodeInfo(ret int, err_msg string) string {
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
