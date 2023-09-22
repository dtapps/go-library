package nldyp

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type PartnerData4LockSeatResponse struct {
	Status int `json:"status"`
	Data   []struct {
		OrderId          string   `json:"orderId"`   // 订单 id
		OrderNo          string   `json:"orderNo"`   // 系统商锁座订单号
		SerialNum        string   `json:"serialNum"` // 锁座流水号
		Direct           int      `json:"direct"`    // 是否直签
		BackTicketConfig []string `json:"backTicketConfig"`
	} `json:"data"`
	Content string `json:"content"`
}

type PartnerData4LockSeatResult struct {
	Result PartnerData4LockSeatResponse // 结果
	Body   []byte                       // 内容
	Http   gorequest.Response           // 请求
}

func newPartnerData4LockSeatResult(result PartnerData4LockSeatResponse, body []byte, http gorequest.Response) *PartnerData4LockSeatResult {
	return &PartnerData4LockSeatResult{Result: result, Body: body, Http: http}
}

// PartnerData4LockSeat 锁座（支持多座区下单）
// https://docs.apipost.cn/preview/fa101f4865dc783f/66e7c2e894fda4a6?target_id=ac7d2885-c575-4efa-8438-03175f8978a9
func (c *Client) PartnerData4LockSeat(ctx context.Context, notMustParams ...*gorequest.Params) (*PartnerData4LockSeatResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/partner/data4/lockSeat", params)
	if err != nil {
		return newPartnerData4LockSeatResult(PartnerData4LockSeatResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response PartnerData4LockSeatResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newPartnerData4LockSeatResult(response, request.ResponseBody, request), err
}
