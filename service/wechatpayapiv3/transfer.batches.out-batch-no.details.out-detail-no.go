package wechatpayapiv3

import (
	"context"
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
	"time"
)

type TransferBatchesOutBatchNoDetailsOutDetailResponse struct {
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

type TransferBatchesOutBatchNoDetailsOutDetailResult struct {
	Result TransferBatchesOutBatchNoDetailsOutDetailResponse // 结果
	Body   []byte                                            // 内容
	Http   gorequest.Response                                // 请求
	Err    error                                             // 错误
}

func newTransferBatchesOutBatchNoDetailsOutDetailResult(result TransferBatchesOutBatchNoDetailsOutDetailResponse, body []byte, http gorequest.Response, err error) *TransferBatchesOutBatchNoDetailsOutDetailResult {
	return &TransferBatchesOutBatchNoDetailsOutDetailResult{Result: result, Body: body, Http: http, Err: err}
}

// TransferBatchesOutBatchNoDetailsOutDetail 商家明细单号查询明细单API
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter4_3_6.shtml
func (c *Client) TransferBatchesOutBatchNoDetailsOutDetail(ctx context.Context, outBatchNo string, outDetailNo string) *TransferBatchesOutBatchNoDetailsOutDetailResult {
	// 参数
	params := gorequest.NewParams()
	// 请求
	request, err := c.request(ctx, apiUrl+"/v3/transfer/batches/out-batch-no/"+outBatchNo+"details/out-detail-no//"+outDetailNo, params, http.MethodGet, false)
	if err != nil {
		return newTransferBatchesOutBatchNoDetailsOutDetailResult(TransferBatchesOutBatchNoDetailsOutDetailResponse{}, request.ResponseBody, request, err)
	}
	// 定义
	var response TransferBatchesOutBatchNoDetailsOutDetailResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newTransferBatchesOutBatchNoDetailsOutDetailResult(response, request.ResponseBody, request, err)
}
