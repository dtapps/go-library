package nldyp

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type PartnerCommonGetSurplusMoneyResponse struct {
	Status int `json:"status"`
	Data   []struct {
		SurplusMoney     float64 `json:"surplusMoney"`     // 余额
		WithholdingMoney float64 `json:"withholdingMoney"` // 预扣金额
		AvailableMoney   float64 `json:"AvailableMoney"`   // 可用余额
	} `json:"data"`
	Content string `json:"content"`
}

type PartnerCommonGetSurplusMoneyResult struct {
	Result PartnerCommonGetSurplusMoneyResponse // 结果
	Body   []byte                               // 内容
	Http   gorequest.Response                   // 请求
	Err    error                                // 错误
}

func newPartnerCommonGetSurplusMoneyResult(result PartnerCommonGetSurplusMoneyResponse, body []byte, http gorequest.Response, err error) *PartnerCommonGetSurplusMoneyResult {
	return &PartnerCommonGetSurplusMoneyResult{Result: result, Body: body, Http: http, Err: err}
}

// PartnerCommonGetSurplusMoney 释放锁座
// https://docs.apipost.cn/preview/fa101f4865dc783f/66e7c2e894fda4a6?target_id=802c4269-60c5-4d61-bda2-ec82cae60930
func (c *Client) PartnerCommonGetSurplusMoney(ctx context.Context, notMustParams ...*gorequest.Params) *PartnerCommonGetSurplusMoneyResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/partner/common/GetSurplusMoney", params)
	// 定义
	var response PartnerCommonGetSurplusMoneyResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newPartnerCommonGetSurplusMoneyResult(response, request.ResponseBody, request, err)
}
