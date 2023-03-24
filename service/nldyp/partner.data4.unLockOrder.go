package nldyp

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type PartnerData4UnLockOrderResponse struct {
	Status int `json:"status"`
	Data   []struct {
		OrderId string `json:"orderId"` // 订单 id
		OrderNo string `json:"orderNo"` // 系统商锁座订单号
	} `json:"data"`
	Content string `json:"content"`
}

type PartnerData4UnLockOrderResult struct {
	Result PartnerData4UnLockOrderResponse // 结果
	Body   []byte                          // 内容
	Http   gorequest.Response              // 请求
	Err    error                           // 错误
}

func newPartnerData4UnLockOrderResult(result PartnerData4UnLockOrderResponse, body []byte, http gorequest.Response, err error) *PartnerData4UnLockOrderResult {
	return &PartnerData4UnLockOrderResult{Result: result, Body: body, Http: http, Err: err}
}

// PartnerData4UnLockOrder 释放锁座
// https://docs.apipost.cn/preview/fa101f4865dc783f/66e7c2e894fda4a6?target_id=43074323-fd3d-4c14-9a17-a447101b410f
func (c *Client) PartnerData4UnLockOrder(ctx context.Context, notMustParams ...gorequest.Params) *PartnerData4UnLockOrderResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/partner/data4/unLockOrder", params)
	// 定义
	var response PartnerData4UnLockOrderResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newPartnerData4UnLockOrderResult(response, request.ResponseBody, request, err)
}
