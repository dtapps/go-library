package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

// GetShareRatio 查询分账比例
// https://developers.weixin.qq.com/doc/oplatform/openApi/ams/percentage/api_getshareratio.html
func (c *Client) GetShareRatio(ctx context.Context, notMustParams ...*gorequest.Params) (response APIRetResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("appid", c.GetAuthorizerAppid())

	// 请求
	err = c.request(ctx, "wxa/getdefaultamsinfo?action=get_share_ratio&access_token="+c.GetComponentAccessToken(), params, http.MethodPost, &response)
	return
}

// ErrcodeInfo 错误描述
func GetGetShareRatioErrcodeInfo(ret int, err_msg string) string {
	switch ret {
	case -202:
		return "内部错误"
	case 1700:
		return "参数错误"
	case 1701:
		return "参数错误"
	case 1735:
		return "商户未完成协议签署流程"
	case 1737:
		return "操作过快"
	case 2056:
		return "服务商未在变现专区开通账户"
	default:
		return err_msg
	}
}
