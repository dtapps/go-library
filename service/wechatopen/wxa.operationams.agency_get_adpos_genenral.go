package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WxaOperationamsAgencyGetAdposGenenralResponse struct {
	Ret    int    `json:"ret"`
	ErrMsg string `json:"err_msg,omitempty"`
	List   []struct {
		SlotId          int64   `json:"slot_id"`
		AdSlot          string  `json:"ad_slot"`
		Date            string  `json:"date"`
		ReqSuccCount    float64 `json:"req_succ_count"`
		ExposureCount   float64 `json:"exposure_count"`
		ExposureRate    float64 `json:"exposure_rate"`
		ClickCount      float64 `json:"click_count"`
		ClickRate       float64 `json:"click_rate"`
		PublisherIncome float64 `json:"publisher_income"`
		Ecpm            float64 `json:"ecpm"`
	} `json:"list"`
	Summary struct {
		ReqSuccCount    float64 `json:"req_succ_count"`
		ExposureCount   float64 `json:"exposure_count"`
		ExposureRate    float64 `json:"exposure_rate"`
		ClickCount      float64 `json:"click_count"`
		ClickRate       float64 `json:"click_rate"`
		PublisherIncome float64 `json:"publisher_income"`
		Ecpm            float64 `json:"ecpm"`
	} `json:"summary"`
	TotalNum int `json:"total_num"`
}

type WxaOperationamsAgencyGetAdposGenenralResult struct {
	Result WxaOperationamsAgencyGetAdposGenenralResponse // 结果
	Body   []byte                                        // 内容
	Http   gorequest.Response                            // 请求
}

func newWxaOperationamsAgencyGetAdposGenenralResult(result WxaOperationamsAgencyGetAdposGenenralResponse, body []byte, http gorequest.Response) *WxaOperationamsAgencyGetAdposGenenralResult {
	return &WxaOperationamsAgencyGetAdposGenenralResult{Result: result, Body: body, Http: http}
}

// WxaOperationamsAgencyGetAdposGenenral
// 获取小程序广告汇总数据
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/ams/ad-data/GetAdposGenenral.html
func (c *Client) WxaOperationamsAgencyGetAdposGenenral(ctx context.Context, authorizerAccessToken string, page, pageSize int64, startDate, endDate, adSlot string, notMustParams ...gorequest.Params) (*WxaOperationamsAgencyGetAdposGenenralResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("page", page)
	params.Set("page_size", pageSize)
	params.Set("start_date", startDate)
	params.Set("end_date", endDate)
	if adSlot != "" {
		params.Set("ad_slot", adSlot)
	}
	// 请求
	request, err := c.request(ctx, apiUrl+"/wxa/operationams?action=agency_get_adpos_genenral&access_token="+authorizerAccessToken, params, http.MethodPost)
	if err != nil {
		return newWxaOperationamsAgencyGetAdposGenenralResult(WxaOperationamsAgencyGetAdposGenenralResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WxaOperationamsAgencyGetAdposGenenralResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWxaOperationamsAgencyGetAdposGenenralResult(response, request.ResponseBody, request), err
}
