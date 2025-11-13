package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type GetCustomShareRatioResponse struct {
	APIRetResponse     // 错误
	ShareRatio     int `json:"share_ratio"`
}

// GetCustomShareRatio 查询自定义分账比例
// https://developers.weixin.qq.com/doc/oplatform/openApi/ams/percentage/api_getcustomshareratio.html
func (c *Client) GetCustomShareRatio(ctx context.Context, notMustParams ...*gorequest.Params) (response GetCustomShareRatioResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("appid", c.GetAuthorizerAppid())

	// 请求
	err = c.request(ctx, "wxa/getdefaultamsinfo?action=agency_get_custom_share_ratio&access_token="+c.GetComponentAccessToken(), params, http.MethodPost, &response)
	return
}

// ErrcodeInfo 错误描述
func GetGetCustomShareRatioErrcodeInfo(ret int, err_msg string) string {
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
	case 2061:
		return "不存在为该appid设置的个性化分成比例"
	default:
		return err_msg
	}
}
