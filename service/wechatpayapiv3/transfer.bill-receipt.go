package wechatpayapiv3

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
	"time"
)

type TransferBillReceiptResponse struct {
	OutBatchNo      string    `json:"out_batch_no"`     // 商户系统内部的商家批次单号，在商户系统内部唯一。需要电子回单的批次单号
	SignatureNo     string    `json:"signature_no"`     // 电子回单申请单号，申请单据的唯一标识
	SignatureStatus string    `json:"signature_status"` // ACCEPTED:已受理，电子签章已受理成功 FINISHED:已完成。电子签章已处理完成
	HashType        string    `json:"hash_type"`        // 电子回单文件的hash方法
	HashValue       string    `json:"hash_value"`       // 电子回单文件的hash值，用于下载之后验证文件的完整、正确性
	DownloadUrl     string    `json:"download_url"`     // 电子回单文件的下载地址
	CreateTime      time.Time `json:"create_time"`      // 电子签章单创建时间，按照使用rfc3339所定义的格式，格式为YYYY-MM-DDThh:mm:ss+TIMEZONE
	UpdateTime      time.Time `json:"update_time"`      // 电子签章单最近一次状态变更的时间，按照使用rfc3339所定义的格式，格式为YYYY-MM-DDThh:mm:ss+TIMEZONE
}

type TransferBillReceiptResult struct {
	Result TransferBillReceiptResponse // 结果
	Body   []byte                      // 内容
	Http   gorequest.Response          // 请求
}

func newTransferBillReceiptResult(result TransferBillReceiptResponse, body []byte, http gorequest.Response) *TransferBillReceiptResult {
	return &TransferBillReceiptResult{Result: result, Body: body, Http: http}
}

// TransferBillReceipt 转账账单电子回单申请受理接口
// https://pay.weixin.qq.com/docs/merchant/apis/batch-transfer-to-balance/electronic-signature/create-electronic-signature.html
func (c *Client) TransferBillReceipt(ctx context.Context, notMustParams ...*gorequest.Params) (*TransferBillReceiptResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/v3/transfer/bill-receipt", params, http.MethodPost, false)
	if err != nil {
		return newTransferBillReceiptResult(TransferBillReceiptResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response TransferBillReceiptResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newTransferBillReceiptResult(response, request.ResponseBody, request), err
}
