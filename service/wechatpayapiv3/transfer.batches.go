package wechatpayapiv3

import (
	"context"
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type TransferBatchesResponse struct {
	OutBatchNo string `json:"out_batch_no"` // 商家批次单号
	BatchId    string `json:"batch_id"`     // 微信批次单号
	CreateTime string `json:"create_time"`  // 批次创建时间
}

type TransferBatchesResult struct {
	Result TransferBatchesResponse // 结果
	Body   []byte                  // 内容
	Http   gorequest.Response      // 请求
	Err    error                   // 错误
}

func newTransferBatchesResult(result TransferBatchesResponse, body []byte, http gorequest.Response, err error) *TransferBatchesResult {
	return &TransferBatchesResult{Result: result, Body: body, Http: http, Err: err}
}

// TransferBatches 发起商家转账API
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter4_3_1.shtml
func (c *Client) TransferBatches(ctx context.Context, notMustParams ...gorequest.Params) *TransferBatchesResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("appid", c.GetAppId())
	// 请求
	request, err := c.request(ctx, apiUrl+"/v3/transfer/batches", params, http.MethodPost, false)
	if err != nil {
		return newTransferBatchesResult(TransferBatchesResponse{}, request.ResponseBody, request, err)
	}
	// 定义
	var response TransferBatchesResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newTransferBatchesResult(response, request.ResponseBody, request, err)
}
