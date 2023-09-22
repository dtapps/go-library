package nldyp

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type PartnerData4GetOrderDataResponse struct {
	Status  int    `json:"status"`
	Content string `json:"content"`
}

type PartnerData4GetOrderDataResult struct {
	Result PartnerData4GetOrderDataResponse // 结果
	Body   []byte                           // 内容
	Http   gorequest.Response               // 请求
}

func newPartnerData4GetOrderDataResult(result PartnerData4GetOrderDataResponse, body []byte, http gorequest.Response) *PartnerData4GetOrderDataResult {
	return &PartnerData4GetOrderDataResult{Result: result, Body: body, Http: http}
}

// PartnerData4GetOrderData 15分钟出票模式
// https://docs.apipost.cn/preview/fa101f4865dc783f/66e7c2e894fda4a6?target_id=98dfc19f-6f76-4018-8de4-86cdeea4fcab
func (c *Client) PartnerData4GetOrderData(ctx context.Context, notMustParams ...*gorequest.Params) (*PartnerData4GetOrderDataResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/partner/data4/GetOrderData", params)
	if err != nil {
		return newPartnerData4GetOrderDataResult(PartnerData4GetOrderDataResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response PartnerData4GetOrderDataResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newPartnerData4GetOrderDataResult(response, request.ResponseBody, request), err
}
