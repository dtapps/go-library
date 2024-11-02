package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type WxaGetDefaultamsInfoAgencyGetCustomShareRatioResponse struct {
	Ret        int    `json:"ret"`
	ErrMsg     string `json:"err_msg"`
	ShareRatio int    `json:"share_ratio"`
}

type WxaGetDefaultamsInfoAgencyGetCustomShareRatioResult struct {
	Result WxaGetDefaultamsInfoAgencyGetCustomShareRatioResponse // 结果
	Body   []byte                                                // 内容
	Http   gorequest.Response                                    // 请求
}

func newWxaGetDefaultamsInfoAgencyGetCustomShareRatioResult(result WxaGetDefaultamsInfoAgencyGetCustomShareRatioResponse, body []byte, http gorequest.Response) *WxaGetDefaultamsInfoAgencyGetCustomShareRatioResult {
	return &WxaGetDefaultamsInfoAgencyGetCustomShareRatioResult{Result: result, Body: body, Http: http}
}

// WxaGetDefaultamsInfoAgencyGetCustomShareRatio 查询自定义分账比例
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/ams/percentage/GetCustomShareRatio.html
func (c *Client) WxaGetDefaultamsInfoAgencyGetCustomShareRatio(ctx context.Context, authorizerAppid, authorizerAccessToken string, notMustParams ...gorequest.Params) (*WxaGetDefaultamsInfoAgencyGetCustomShareRatioResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("appid", authorizerAppid)

	// 请求
	var response WxaGetDefaultamsInfoAgencyGetCustomShareRatioResponse
	request, err := c.request(ctx, "wxa/getdefaultamsinfo?action=agency_get_custom_share_ratio&access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newWxaGetDefaultamsInfoAgencyGetCustomShareRatioResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *WxaGetDefaultamsInfoAgencyGetCustomShareRatioResult) ErrcodeInfo() string {
	switch resp.Result.Ret {
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
		return resp.Result.ErrMsg
	}
}
