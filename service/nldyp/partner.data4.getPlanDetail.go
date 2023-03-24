package nldyp

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type PartnerData4GetPlanDetailResponse struct {
	Code int `json:"code"`
	Data []struct {
		FeatureAppNo  string `json:"featureAppNo"` // 排期编码
		CinemaCode    int    `json:"cinemaCode"`   // 影城编码
		SourceFilmNo  string `json:"sourceFilmNo"`
		FilmNo        string `json:"filmNo"`        // 影片编码
		FilmName      string `json:"filmName"`      // 影片名称
		HallNo        string `json:"hallNo"`        // 影厅编码
		HallName      string `json:"hallName"`      // 影厅名称
		StartTime     int    `json:"startTime"`     // 放映时间
		CopyType      string `json:"copyType"`      // 影片制式
		CopyLanguage  string `json:"copyLanguage"`  // 语言
		TotalTime     string `json:"totalTime"`     // 时长
		ListingPrice  string `json:"listingPrice"`  // 挂牌价
		TicketPrice   string `json:"ticketPrice"`   // 票价
		ServiceAddFee string `json:"serviceAddFee"` // 服务费下限
		LowestPrice   string `json:"lowestPrice"`   // 最低保护价
		Thresholds    string `json:"thresholds"`    // 服务费上限
		Areas         string `json:"areas"`         // 座区属性
		MarketPrice   string `json:"marketPrice"`   // 市场参考价（无座位区间时，下特价票使用该价格，有座位区间则使用座位区间价）
	} `json:"data"`
	Content string `json:"content"`
}

type PartnerData4GetPlanDetailResult struct {
	Result PartnerData4GetPlanDetailResponse // 结果
	Body   []byte                            // 内容
	Http   gorequest.Response                // 请求
	Err    error                             // 错误
}

func newPartnerData4GetPlanDetailResult(result PartnerData4GetPlanDetailResponse, body []byte, http gorequest.Response, err error) *PartnerData4GetPlanDetailResult {
	return &PartnerData4GetPlanDetailResult{Result: result, Body: body, Http: http, Err: err}
}

// PartnerData4GetPlanDetail 获取影城单个排期详情
// https://docs.apipost.cn/preview/fa101f4865dc783f/66e7c2e894fda4a6?target_id=05639c5c-080f-43f4-b94d-e4b8bb14130e
func (c *Client) PartnerData4GetPlanDetail(ctx context.Context, cinemaId int, featureAppNo string) *PartnerData4GetPlanDetailResult {
	// 参数
	params := gorequest.NewParams()
	params.Set("cinemaId", cinemaId)
	params.Set("featureAppNo", featureAppNo)
	// 请求
	request, err := c.request(ctx, apiUrl+"/partner/data4/getPlanDetail", params)
	// 定义
	var response PartnerData4GetPlanDetailResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newPartnerData4GetPlanDetailResult(response, request.ResponseBody, request, err)
}
