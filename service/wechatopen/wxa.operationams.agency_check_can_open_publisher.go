package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type AgencyCheckCanOpenPublisherResponse struct {
	APIRetResponse     // 错误
	Status         int `json:"status"`
}

// AgencyCheckCanOpenPublisher 检测是否能开通流量主
// https://developers.weixin.qq.com/doc/oplatform/openApi/ams/open/api_agencycheckcanopenpublisher.html
func (c *Client) AgencyCheckCanOpenPublisher(ctx context.Context, notMustParams ...*gorequest.Params) (response AgencyCheckCanOpenPublisherResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "wxa/operationams?action=agency_check_can_open_publisher&access_token="+c.GetAuthorizerAccessToken(), params, http.MethodPost, &response)
	return
}

// ErrcodeInfo 错误描述
func GetAgencyCheckCanOpenPublisherErrcodeInfo(ret int, err_msg string) string {
	switch ret {
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
