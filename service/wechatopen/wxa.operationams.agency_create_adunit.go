package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WxaOperationamsAgencyCreateAdunitResponse struct {
	Ret      int    `json:"ret"`
	ErrMsg   string `json:"err_msg,omitempty"`
	AdUnitId string `json:"ad_unit_id"`
}

type WxaOperationamsAgencyCreateAdunitResult struct {
	Result WxaOperationamsAgencyCreateAdunitResponse // 结果
	Body   []byte                                    // 内容
	Http   gorequest.Response                        // 请求
}

func newWxaOperationamsAgencyCreateAdunitResult(result WxaOperationamsAgencyCreateAdunitResponse, body []byte, http gorequest.Response) *WxaOperationamsAgencyCreateAdunitResult {
	return &WxaOperationamsAgencyCreateAdunitResult{Result: result, Body: body, Http: http}
}

// WxaOperationamsAgencyCreateAdunit
// 创建广告单元
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/ams/ad-mgnt/AgencyCreateAdunit.html
func (c *Client) WxaOperationamsAgencyCreateAdunit(ctx context.Context, authorizerAccessToken string, name, Type string, videoDurationMin, videoDurationMax int64, notMustParams ...gorequest.Params) (*WxaOperationamsAgencyCreateAdunitResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("name", name)
	params.Set("type", Type)
	if videoDurationMin > 0 {
		params.Set("video_duration_min", videoDurationMin)
	}
	if videoDurationMax > 0 {
		params.Set("video_duration_max", videoDurationMax)
	}
	// 请求
	request, err := c.request(ctx, apiUrl+"/wxa/operationams?action=agency_create_adunit&access_token="+authorizerAccessToken, params, http.MethodPost)
	if err != nil {
		return newWxaOperationamsAgencyCreateAdunitResult(WxaOperationamsAgencyCreateAdunitResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WxaOperationamsAgencyCreateAdunitResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWxaOperationamsAgencyCreateAdunitResult(response, request.ResponseBody, request), err
}
