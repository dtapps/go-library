package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

// WxaOperationamsAgencyCreatePublisher
// 开通流量主
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/ams/open/AgencyCreatePublisher.html
func (c *Client) WxaOperationamsAgencyCreatePublisher(ctx context.Context, notMustParams ...*gorequest.Params) (response APIRetResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "wxa/operationams?action=agency_create_publisher&access_token="+c.GetAuthorizerAccessToken(), params, http.MethodPost, &response)
	return
}

// ErrcodeInfo 错误描述
func GetWxaOperationamsAgencyCreatePublisherErrcodeInfo(ret int, err_msg string) string {
	switch ret {
	case 1700:
		return "参数错误"
	case 1701:
		return "参数错误"
	case 1735:
		return "商户未完成协议签署流程"
	case 1737:
		return "操作过快"
	case 1807:
		return "无效流量主"
	case 2009:
		return "无效流量主"
	case 2021:
		return "已开通流量主"
	case 2056:
		return "服务商未在变现专区开通账户"
	case 2013:
		return "未满足开通流量主门槛（1000个独立访问用户UV）"
	default:
		return err_msg
	}
}
