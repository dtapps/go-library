package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WxaSetDefaultamsInfoAgencySetCustomShareRatioResponse struct {
	Ret    int    `json:"ret"`
	ErrMsg string `json:"err_msg"`
}

type WxaSetDefaultamsInfoAgencySetCustomShareRatioResult struct {
	Result WxaSetDefaultamsInfoAgencySetCustomShareRatioResponse // 结果
	Body   []byte                                                // 内容
	Http   gorequest.Response                                    // 请求
}

func newWxaSetDefaultamsInfoAgencySetCustomShareRatioResult(result WxaSetDefaultamsInfoAgencySetCustomShareRatioResponse, body []byte, http gorequest.Response) *WxaSetDefaultamsInfoAgencySetCustomShareRatioResult {
	return &WxaSetDefaultamsInfoAgencySetCustomShareRatioResult{Result: result, Body: body, Http: http}
}

// WxaSetDefaultamsInfoAgencySetCustomShareRatio 设置自定义分账比例
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/ams/percentage/SetCustomShareRatio.html
func (c *Client) WxaSetDefaultamsInfoAgencySetCustomShareRatio(ctx context.Context, shareRatio int64, notMustParams ...gorequest.Params) (*WxaSetDefaultamsInfoAgencySetCustomShareRatioResult, error) {
	// 检查
	if err := c.checkAuthorizerConfig(ctx); err != nil {
		return newWxaSetDefaultamsInfoAgencySetCustomShareRatioResult(WxaSetDefaultamsInfoAgencySetCustomShareRatioResponse{}, []byte{}, gorequest.Response{}), err
	}
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	if shareRatio > 0 {
		params.Set("share_ratio", shareRatio)
	}
	// 请求
	request, err := c.request(ctx, apiUrl+"/wxa/setdefaultamsinfo?action=agency_set_custom_share_ratio&access_token="+c.GetAuthorizerAccessToken(ctx), params, http.MethodPost)
	if err != nil {
		return newWxaSetDefaultamsInfoAgencySetCustomShareRatioResult(WxaSetDefaultamsInfoAgencySetCustomShareRatioResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WxaSetDefaultamsInfoAgencySetCustomShareRatioResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWxaSetDefaultamsInfoAgencySetCustomShareRatioResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *WxaSetDefaultamsInfoAgencySetCustomShareRatioResult) ErrcodeInfo() string {
	switch resp.Result.Ret {
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
	}
	return "系统繁忙"
}
