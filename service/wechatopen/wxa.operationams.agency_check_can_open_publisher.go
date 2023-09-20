package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WxaOperationamsAgencyCheckCanOpenPublisherResponse struct {
	Ret    int    `json:"ret"`
	ErrMsg string `json:"err_msg"`
	Status int    `json:"status"`
}

type WxaOperationamsAgencyCheckCanOpenPublisherResult struct {
	Result WxaOperationamsAgencyCheckCanOpenPublisherResponse // 结果
	Body   []byte                                             // 内容
	Http   gorequest.Response                                 // 请求
}

func newWxaOperationamsAgencyCheckCanOpenPublisherResult(result WxaOperationamsAgencyCheckCanOpenPublisherResponse, body []byte, http gorequest.Response) *WxaOperationamsAgencyCheckCanOpenPublisherResult {
	return &WxaOperationamsAgencyCheckCanOpenPublisherResult{Result: result, Body: body, Http: http}
}

// WxaOperationamsAgencyCheckCanOpenPublisher 检测是否能开通流量主
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/ams/open/AgencyCheckCanOpenPublisher.html
func (c *Client) WxaOperationamsAgencyCheckCanOpenPublisher(ctx context.Context, authorizerAccessToken string, notMustParams ...*gorequest.Params) (*WxaOperationamsAgencyCheckCanOpenPublisherResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/wxa/operationams?action=agency_check_can_open_publisher&access_token="+authorizerAccessToken, params, http.MethodPost)
	if err != nil {
		return newWxaOperationamsAgencyCheckCanOpenPublisherResult(WxaOperationamsAgencyCheckCanOpenPublisherResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WxaOperationamsAgencyCheckCanOpenPublisherResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWxaOperationamsAgencyCheckCanOpenPublisherResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *WxaOperationamsAgencyCheckCanOpenPublisherResult) ErrcodeInfo() string {
	switch resp.Result.Ret {
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
	}
	return "系统繁忙"
}
