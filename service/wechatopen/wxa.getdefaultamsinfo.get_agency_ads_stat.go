package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type WxaGetdefaultamsinfoGetAgencyAdsStatResponse struct {
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
		Income          float64 `json:"income"`
		Ecpm            float64 `json:"ecpm"`
		AgencyIncome    float64 `json:"agency_income"`
		PublisherIncome float64 `json:"publisher_income"`
		PublisherAppid  string  `json:"publisher_appid"`
	} `json:"list"`
	Summary struct {
		ReqSuccCount    float64 `json:"req_succ_count"`
		ExposureCount   float64 `json:"exposure_count"`
		ExposureRate    float64 `json:"exposure_rate"`
		ClickCount      float64 `json:"click_count"`
		ClickRate       float64 `json:"click_rate"`
		Income          float64 `json:"income"`
		Ecpm            float64 `json:"ecpm"`
		ExposureUv      float64 `json:"exposure_uv"`
		OpenUv          float64 `json:"open_uv"`
		PublisherIncome float64 `json:"publisher_income"`
		AgencyIncome    float64 `json:"agency_income"`
	} `json:"summary"`
	TotalNum int `json:"total_num"`
}

type WxaGetdefaultamsinfoGetAgencyAdsStatResult struct {
	Result WxaGetdefaultamsinfoGetAgencyAdsStatResponse // 结果
	Body   []byte                                       // 内容
	Http   gorequest.Response                           // 请求
}

func newWxaGetdefaultamsinfoGetAgencyAdsStatResult(result WxaGetdefaultamsinfoGetAgencyAdsStatResponse, body []byte, http gorequest.Response) *WxaGetdefaultamsinfoGetAgencyAdsStatResult {
	return &WxaGetdefaultamsinfoGetAgencyAdsStatResult{Result: result, Body: body, Http: http}
}

// WxaGetdefaultamsinfoGetAgencyAdsStat
// 获取服务商广告汇总数据
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/ams/ad-data/GetAgencyAdsStat.html
func (c *Client) WxaGetdefaultamsinfoGetAgencyAdsStat(ctx context.Context, authorizerAccessToken string, page, pageSize int64, startDate, endDate, adSlot string, notMustParams ...gorequest.Params) (*WxaGetdefaultamsinfoGetAgencyAdsStatResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "wxa/getdefaultamsinfo")
	defer span.End()

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
	var response WxaGetdefaultamsinfoGetAgencyAdsStatResponse
	request, err := c.request(ctx, span, "wxa/getdefaultamsinfo?action=get_agency_ads_stat&access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return newWxaGetdefaultamsinfoGetAgencyAdsStatResult(response, request.ResponseBody, request), err
}
