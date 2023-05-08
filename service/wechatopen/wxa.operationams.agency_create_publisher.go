package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WxaOperationamsAgencyCreatePublisherResponse struct {
	Ret    int    `json:"ret"`
	ErrMsg string `json:"err_msg"`
}

type WxaOperationamsAgencyCreatePublisherResult struct {
	Result WxaOperationamsAgencyCreatePublisherResponse // 结果
	Body   []byte                                       // 内容
	Http   gorequest.Response                           // 请求
}

func newWxaOperationamsAgencyCreatePublisherResult(result WxaOperationamsAgencyCreatePublisherResponse, body []byte, http gorequest.Response) *WxaOperationamsAgencyCreatePublisherResult {
	return &WxaOperationamsAgencyCreatePublisherResult{Result: result, Body: body, Http: http}
}

// WxaOperationamsAgencyCreatePublisher 检测是否能开通流量主
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/ams/open/AgencyCheckCanOpenPublisher.html
func (c *Client) WxaOperationamsAgencyCreatePublisher(ctx context.Context, notMustParams ...gorequest.Params) (*WxaOperationamsAgencyCreatePublisherResult, error) {
	// 检查
	if err := c.checkAuthorizerConfig(ctx); err != nil {
		return newWxaOperationamsAgencyCreatePublisherResult(WxaOperationamsAgencyCreatePublisherResponse{}, []byte{}, gorequest.Response{}), err
	}
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/wxa/operationams?action=agency_check_can_open_publisher&access_token="+GetAuthorizerAccessToken(ctx, c), params, http.MethodPost)
	if err != nil {
		return newWxaOperationamsAgencyCreatePublisherResult(WxaOperationamsAgencyCreatePublisherResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WxaOperationamsAgencyCreatePublisherResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWxaOperationamsAgencyCreatePublisherResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *WxaOperationamsAgencyCreatePublisherResult) ErrcodeInfo() string {
	switch resp.Result.Ret {
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
	}
	return "系统繁忙"
}
