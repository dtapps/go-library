package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WxaSetDefaultamsInfoSetShareRatioResponse struct {
	Ret    int    `json:"ret"`
	ErrMsg string `json:"err_msg"`
}

type WxaSetDefaultamsInfoSetShareRatioResult struct {
	Result WxaSetDefaultamsInfoSetShareRatioResponse // 结果
	Body   []byte                                    // 内容
	Http   gorequest.Response                        // 请求
}

func newWxaSetDefaultamsInfoSetShareRatioResult(result WxaSetDefaultamsInfoSetShareRatioResponse, body []byte, http gorequest.Response) *WxaSetDefaultamsInfoSetShareRatioResult {
	return &WxaSetDefaultamsInfoSetShareRatioResult{Result: result, Body: body, Http: http}
}

// WxaSetDefaultamsInfoSetShareRatio 设置默认分账比例
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/ams/percentage/SetShareRatio.html
func (c *Client) WxaSetDefaultamsInfoSetShareRatio(ctx context.Context, shareRatio int64, notMustParams ...gorequest.Params) (*WxaSetDefaultamsInfoSetShareRatioResult, error) {
	// 检查
	if err := c.checkAuthorizerConfig(ctx); err != nil {
		return newWxaSetDefaultamsInfoSetShareRatioResult(WxaSetDefaultamsInfoSetShareRatioResponse{}, []byte{}, gorequest.Response{}), err
	}
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	if shareRatio > 0 {
		params.Set("share_ratio", shareRatio)
	}
	// 请求
	request, err := c.request(ctx, apiUrl+"/wxa/setdefaultamsinfo?action=set_share_ratio&access_token="+GetAuthorizerAccessToken(ctx, c), params, http.MethodPost)
	if err != nil {
		return newWxaSetDefaultamsInfoSetShareRatioResult(WxaSetDefaultamsInfoSetShareRatioResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WxaSetDefaultamsInfoSetShareRatioResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
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
	}
	return "系统繁忙"
}
