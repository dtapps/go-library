package nldyp

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type PartnerData4GetPlanSeatResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		SeatState     int    `json:"seatState"` // 状 态 (-1 不 可 售 0-可售)
		SectionCode   string `json:"sectionCode"`
		SeatNo        string `json:"seatNo"`        // 座位编号
		GraphRow      int    `json:"graphRow"`      // 逻辑坐标行（绘图）
		GraphCol      int    `json:"graphCol"`      // 逻辑坐标列（绘图）
		SeatRow       string `json:"seatRow"`       // 物理座位行号
		SeatCol       string `json:"seatCol"`       // 物理座位列号
		SeatPieceNo   string `json:"seatPieceNo"`   // 连座编号，带相同编 号的必须同时锁座
		SeatPieceName string `json:"seatPieceName"` // 座位名称
		AreaId        string `json:"areaId"`        // 座区 Id，该字段为 空表示未设置座区
		SeatType      string `json:"seatType"`      // 座 位 类 型 N: 普 通 座,L:情侣座首座,M: 三连中间座，R:情侣 座次座
		SeatFlag      int    `json:"seatFlag"`
	} `json:"data"`
}

type PartnerData4GetPlanSeatResult struct {
	Result PartnerData4GetPlanSeatResponse // 结果
	Body   []byte                          // 内容
	Http   gorequest.Response              // 请求
}

func newPartnerData4GetPlanSeatResult(result PartnerData4GetPlanSeatResponse, body []byte, http gorequest.Response) *PartnerData4GetPlanSeatResult {
	return &PartnerData4GetPlanSeatResult{Result: result, Body: body, Http: http}
}

// PartnerData4GetPlanSeat 获取实时座位图
// https://docs.apipost.cn/preview/fa101f4865dc783f/66e7c2e894fda4a6?target_id=b245efe6-f728-450a-92f4-a93669c1d555
func (c *Client) PartnerData4GetPlanSeat(ctx context.Context, notMustParams ...*gorequest.Params) (*PartnerData4GetPlanSeatResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/partner/data4/getPlanSeat", params)
	if err != nil {
		return newPartnerData4GetPlanSeatResult(PartnerData4GetPlanSeatResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response PartnerData4GetPlanSeatResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newPartnerData4GetPlanSeatResult(response, request.ResponseBody, request), err
}
