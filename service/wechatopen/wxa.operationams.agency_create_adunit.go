package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type WxaOperationamsAgencyCreateAdunitResponse struct {
	APIRetResponse        // 错误
	AdUnitId       string `json:"ad_unit_id"`
}

// WxaOperationamsAgencyCreateAdunit
// 创建广告单元
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/ams/ad-mgnt/AgencyCreateAdunit.html
func (c *Client) WxaOperationamsAgencyCreateAdunit(ctx context.Context, name, Type string, videoDurationMin, videoDurationMax int64, notMustParams ...*gorequest.Params) (response WxaOperationamsAgencyCreateAdunitResponse, err error) {

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
	err = c.request(ctx, "wxa/operationams?action=agency_create_adunit&access_token="+c.GetAuthorizerAccessToken(), params, http.MethodPost, &response)
	return
}
