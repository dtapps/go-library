package wechatpayapiv3

import (
	"context"
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
	"time"
)

type TransferBatchesBatchIdDetailsDetailIdResponse struct {
	Mchid          string    `json:"mchid"`           // 商户号
	OutBatchNo     string    `json:"out_batch_no"`    // 商家批次单号
	BatchId        string    `json:"batch_id"`        // 微信批次单号
	Appid          string    `json:"appid"`           // 直连商户的appid
	OutDetailNo    string    `json:"out_detail_no"`   // 商家明细单号
	DetailId       string    `json:"detail_id"`       // 微信明细单号
	DetailStatus   string    `json:"detail_status"`   // 明细状态
	TransferAmount int       `json:"transfer_amount"` // 转账金额
	TransferRemark string    `json:"transfer_remark"` // 转账备注
	FailReason     string    `json:"fail_reason"`     // 明细失败原因
	Openid         string    `json:"openid"`          // 用户在直连商户应用下的用户标示
	UserName       string    `json:"user_name"`       // 收款用户姓名
	InitiateTime   time.Time `json:"initiate_time"`   // 转账发起时间
	UpdateTime     time.Time `json:"update_time"`     // 明细更新时间
}

type TransferBatchesBatchIdDetailsDetailIdResult struct {
	Result TransferBatchesBatchIdDetailsDetailIdResponse // 结果
	Body   []byte                                        // 内容
	Http   gorequest.Response                            // 请求
	Err    error                                         // 错误
}

func newTransferBatchesBatchIdDetailsDetailIdResult(result TransferBatchesBatchIdDetailsDetailIdResponse, body []byte, http gorequest.Response, err error) *TransferBatchesBatchIdDetailsDetailIdResult {
	return &TransferBatchesBatchIdDetailsDetailIdResult{Result: result, Body: body, Http: http, Err: err}
}

// TransferBatchesBatchIdDetailsDetailId 微信明细单号查询明细单API
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter4_3_3.shtml
func (c *Client) TransferBatchesBatchIdDetailsDetailId(ctx context.Context, batchId string, detailId string) *TransferBatchesBatchIdDetailsDetailIdResult {
	// 参数
	params := gorequest.NewParams()
	// 请求
	request, err := c.request(ctx, apiUrl+"/v3/transfer/batches/batch-id/"+batchId+"details/detail-id/"+detailId, params, http.MethodGet, false)
	if err != nil {
		return newTransferBatchesBatchIdDetailsDetailIdResult(TransferBatchesBatchIdDetailsDetailIdResponse{}, request.ResponseBody, request, err)
	}
	// 定义
	var response TransferBatchesBatchIdDetailsDetailIdResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newTransferBatchesBatchIdDetailsDetailIdResult(response, request.ResponseBody, request, err)
}
