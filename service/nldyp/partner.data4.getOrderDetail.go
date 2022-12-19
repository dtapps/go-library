package nldyp

import (
	"context"
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
)

type PartnerData4GetOrderDetailResponse struct {
	OrderNo       string  `json:"orderNo"`       // 订单号
	Status        int     `json:"status"`        // 订单状态： 1 出票中 2 已出票 3 出票失败 4 关闭订单 5 系统已出票 6 系统关闭订单
	Mobile        string  `json:"mobile"`        // 手机号码
	BuyTime       int     `json:"buyTime"`       // 下单时间（时间戳）
	FilmName      string  `json:"filmName"`      // 电影名称
	CinemaName    string  `json:"cinemaName"`    // 影院名称
	CinemaAddress string  `json:"cinemaAddress"` // 影院地址
	CopyLanguage  string  `json:"copyLanguage"`  // 影片语言
	MovieType     string  `json:"movieType"`     // 影片类型
	StartTime     int     `json:"startTime"`     // 开场时间（时间戳）
	TicketNum     int     `json:"ticketNum"`     // 购票数量
	HallName      string  `json:"hallName"`      // 影厅名称
	CopyType      string  `json:"copyType"`      // 影厅类型
	CityName      string  `json:"cityName"`      // 城市
	MarketPrice   float64 `json:"marketPrice"`   // 市场价
	CostPrice     float64 `json:"costPrice"`     // 成本价
	BuyType       int     `json:"buyType"`       // 下单类型：0 特惠 1 秒出
}

type PartnerData4GetOrderDetailResult struct {
	Result PartnerData4GetOrderDetailResponse // 结果
	Body   []byte                             // 内容
	Http   gorequest.Response                 // 请求
	Err    error                              // 错误
}

func newPartnerData4GetOrderDetailResult(result PartnerData4GetOrderDetailResponse, body []byte, http gorequest.Response, err error) *PartnerData4GetOrderDetailResult {
	return &PartnerData4GetOrderDetailResult{Result: result, Body: body, Http: http, Err: err}
}

// PartnerData4GetOrderDetail 获取订单详情
// https://docs.apipost.cn/preview/fa101f4865dc783f/66e7c2e894fda4a6?target_id=0888fc18-6ac7-4d37-a1c5-f26e136b381a
func (c *Client) PartnerData4GetOrderDetail(ctx context.Context, orderNo string) *PartnerData4GetOrderDetailResult {
	// 参数
	params := gorequest.NewParams()
	params["orderNo"] = orderNo // 平台订单号
	// 请求
	request, err := c.request(ctx, apiUrl+"/partner/data4/getOrderDetail", params)
	// 定义
	var response PartnerData4GetOrderDetailResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newPartnerData4GetOrderDetailResult(response, request.ResponseBody, request, err)
}
