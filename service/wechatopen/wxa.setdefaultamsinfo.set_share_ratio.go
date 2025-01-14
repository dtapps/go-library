package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type WxaSetDefaultamsInfoSetShareRatioResponse struct {
	Ret    int    `json:"ret"`
	ErrMsg string `json:"err_msg,omitempty"`
}

type WxaSetDefaultamsInfoSetShareRatioResult struct {
	Result WxaSetDefaultamsInfoSetShareRatioResponse // 结果
	Body   []byte                                    // 内容
	Http   gorequest.Response                        // 请求
}

func newWxaSetDefaultamsInfoSetShareRatioResult(result WxaSetDefaultamsInfoSetShareRatioResponse, body []byte, http gorequest.Response) *WxaSetDefaultamsInfoSetShareRatioResult {
	return &WxaSetDefaultamsInfoSetShareRatioResult{Result: result, Body: body, Http: http}
}

// WxaSetDefaultamsInfoSetShareRatio
// 设置默认分账比例
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/ams/percentage/SetShareRatio.html
func (c *Client) WxaSetDefaultamsInfoSetShareRatio(ctx context.Context, authorizerAccessToken string, shareRatio int64, notMustParams ...*gorequest.Params) (*WxaSetDefaultamsInfoSetShareRatioResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("share_ratio", shareRatio)

	// 请求
	var response WxaSetDefaultamsInfoSetShareRatioResponse
	request, err := c.request(ctx, "wxa/setdefaultamsinfo?action=set_share_ratio&access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newWxaSetDefaultamsInfoSetShareRatioResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *WxaSetDefaultamsInfoSetShareRatioResult) ErrcodeInfo() string {
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
	default:
		return resp.Result.ErrMsg
	}
}
