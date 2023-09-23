package wechatpayopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type BillFundFlowBillGetResponse struct {
	DownloadUrl string `json:"download_url"` // 哈希类型
	HashType    string `json:"hash_type"`    // 哈希值
	HashValue   string `json:"hash_value"`   // 账单下载地址
}

type BillFundFlowBillGetResult struct {
	Result BillFundFlowBillGetResponse // 结果
	Body   []byte                      // 内容
	Http   gorequest.Response          // 请求
}

func newBillFundFlowBillGetResult(result BillFundFlowBillGetResponse, body []byte, http gorequest.Response) *BillFundFlowBillGetResult {
	return &BillFundFlowBillGetResult{Result: result, Body: body, Http: http}
}

// BillFundFlowBillGet 申请资金账单API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_4_7.shtml
func (c *Client) BillFundFlowBillGet(ctx context.Context, notMustParams ...*gorequest.Params) (*BillFundFlowBillGetResult, ApiError, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/v3/bill/fundflowbill", params, http.MethodGet)
	if err != nil {
		return newBillFundFlowBillGetResult(BillFundFlowBillGetResponse{}, request.ResponseBody, request), ApiError{}, err
	}
	// 定义
	var response BillFundFlowBillGetResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	// 错误
	var apiError ApiError
	err = gojson.Unmarshal(request.ResponseBody, &apiError)
	return newBillFundFlowBillGetResult(response, request.ResponseBody, request), apiError, err
}
