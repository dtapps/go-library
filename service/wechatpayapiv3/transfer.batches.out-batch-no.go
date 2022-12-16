package wechatpayapiv3

import (
	"context"
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type TransferBatchesOutBatchNoResponse struct {
	TransferBatch struct {
		Mchid         string `json:"mchid"`                  // 商户号
		OutBatchNo    string `json:"out_batch_no"`           // 商家批次单号
		BatchId       string `json:"batch_id"`               // 微信批次单号
		Appid         string `json:"appid"`                  // 直连商户的appid
		BatchStatus   string `json:"batch_status"`           // 批次状态
		BatchType     string `json:"batch_type"`             // 批次类型
		BatchName     string `json:"batch_name"`             // 批次名称
		BatchRemark   string `json:"batch_remark"`           // 批次备注
		CloseReason   string `json:"close_reason,omitempty"` // 批次关闭原因
		TotalAmount   int    `json:"total_amount"`           // 转账总金额
		TotalNum      int    `json:"total_num"`              // 转账总笔数
		CreateTime    string `json:"create_time"`            // 批次创建时间
		UpdateTime    string `json:"update_time"`            // 批次更新时间
		SuccessAmount int    `json:"success_amount"`         // 转账成功金额
		SuccessNum    int    `json:"success_num"`            // 转账成功笔数
		FailAmount    int    `json:"fail_amount"`            // 转账失败金额
		FailNum       int    `json:"fail_num"`               // 转账失败笔数
	} `json:"transfer_batch"` // 转账批次单
	TransferDetailList []struct {
		DetailId     string `json:"detail_id"`     // 微信明细单号
		OutDetailNo  string `json:"out_detail_no"` // 商家明细单号
		DetailStatus string `json:"detail_status"` // 明细状态
	} `json:"transfer_detail_list,omitempty"` // 转账明细单列表
	Offset int `json:"offset"` // 请求资源起始位置
	Limit  int `json:"limit"`  // 最大资源条数
}

type TransferBatchesOutBatchNoResult struct {
	Result TransferBatchesOutBatchNoResponse // 结果
	Body   []byte                            // 内容
	Http   gorequest.Response                // 请求
	Err    error                             // 错误
}

func newTransferBatchesOutBatchNoResult(result TransferBatchesOutBatchNoResponse, body []byte, http gorequest.Response, err error) *TransferBatchesOutBatchNoResult {
	return &TransferBatchesOutBatchNoResult{Result: result, Body: body, Http: http, Err: err}
}

// TransferBatchesOutBatchNo 商家批次单号查询批次单API
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter4_3_5.shtml
func (c *Client) TransferBatchesOutBatchNo(ctx context.Context, outBatchNo string, needQueryDetail bool, offset, limit int, detailStatus string) *TransferBatchesOutBatchNoResult {
	// 参数
	params := gorequest.NewParams()
	params.Set("out_batch_no", outBatchNo)
	params.Set("need_query_detail", needQueryDetail)
	params.Set("offset", offset)
	params.Set("limit", limit)
	if needQueryDetail {
		params.Set("detail_status", detailStatus)
	}
	// 请求
	request, err := c.request(ctx, apiUrl+"/v3/transfer/batches/out-batch-no/"+outBatchNo, params, http.MethodGet, false)
	if err != nil {
		return newTransferBatchesOutBatchNoResult(TransferBatchesOutBatchNoResponse{}, request.ResponseBody, request, err)
	}
	// 定义
	var response TransferBatchesOutBatchNoResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newTransferBatchesOutBatchNoResult(response, request.ResponseBody, request, err)
}
