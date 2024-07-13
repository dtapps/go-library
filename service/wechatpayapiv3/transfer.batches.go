package wechatpayapiv3

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type TransferBatchesResponse struct {
	OutBatchNo string `json:"out_batch_no"` // 商户系统内部的商家批次单号，在商户系统内部唯一
	BatchId    string `json:"batch_id"`     // 微信批次单号，微信商家转账系统返回的唯一标识
	CreateTime string `json:"create_time"`  // 批次受理成功时返回，按照使用rfc3339所定义的格式，格式为YYYY-MM-DDThh:mm:ss+TIMEZONE
}

type TransferBatchesResult struct {
	Result TransferBatchesResponse // 结果
	Body   []byte                  // 内容
	Http   gorequest.Response      // 请求
}

func newTransferBatchesResult(result TransferBatchesResponse, body []byte, http gorequest.Response) *TransferBatchesResult {
	return &TransferBatchesResult{Result: result, Body: body, Http: http}
}

// TransferBatches 发起商家转账
// https://pay.weixin.qq.com/docs/merchant/apis/batch-transfer-to-balance/transfer-batch/initiate-batch-transfer.html
func (c *Client) TransferBatches(ctx context.Context, notMustParams ...gorequest.Params) (*TransferBatchesResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "v3/transfer/batches")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("appid", c.GetAppId())

	// 请求
	var response TransferBatchesResponse
	request, err := c.request(ctx, "v3/transfer/batches", params, http.MethodPost, false, &response)
	return newTransferBatchesResult(response, request.ResponseBody, request), err
}
