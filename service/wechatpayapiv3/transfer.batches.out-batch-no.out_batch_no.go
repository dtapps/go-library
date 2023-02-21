package wechatpayapiv3

import (
	"context"
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type TransferBatchesOutBatchNoOutBatchNoResponse struct {
	TransferBatch struct {
		Mchid           string `json:"mchid"`                  // 微信支付分配的商户号
		OutBatchNo      string `json:"out_batch_no"`           // 商户系统内部的商家批次单号，在商户系统内部唯一
		BatchId         string `json:"batch_id"`               // 微信批次单号，微信商家转账系统返回的唯一标识
		Appid           string `json:"appid"`                  // 申请商户号的appid或商户号绑定的appid（企业号corpid即为此appid）
		BatchStatus     string `json:"batch_status"`           // WAIT_PAY: 待付款确认。需要付款出资商户在商家助手小程序或服务商助手小程序进行付款确认 ACCEPTED:已受理。批次已受理成功，若发起批量转账的30分钟后，转账批次单仍处于该状态，可能原因是商户账户余额不足等。商户可查询账户资金流水，若该笔转账批次单的扣款已经发生，则表示批次已经进入转账中，请再次查单确认 PROCESSING:转账中。已开始处理批次内的转账明细单 FINISHED:已完成。批次内的所有转账明细单都已处理完成 CLOSED:已关闭。可查询具体的批次关闭原因确认
		BatchType       string `json:"batch_type"`             // API:API方式发起 WEB:页面方式发起
		BatchName       string `json:"batch_name"`             // 该笔批量转账的名称
		BatchRemark     string `json:"batch_remark"`           // 转账说明，UTF8编码，最多允许32个字符
		CloseReason     string `json:"close_reason,omitempty"` // 如果批次单状态为“CLOSED”（已关闭），则有关闭原因
		TotalAmount     int    `json:"total_amount"`           // 转账金额单位为“分”
		TotalNum        int    `json:"total_num"`              // 一个转账批次单最多发起三千笔转账
		CreateTime      string `json:"create_time"`            // 批次受理成功时返回，按照使用rfc3339所定义的格式，格式为YYYY-MM-DDThh:mm:ss+TIMEZONE
		UpdateTime      string `json:"update_time"`            // 批次最近一次状态变更的时间，按照使用rfc3339所定义的格式，格式为YYYY-MM-DDThh:mm:ss+TIMEZONE
		SuccessAmount   int    `json:"success_amount"`         // 转账成功的金额，单位为“分”。当批次状态为“PROCESSING”（转账中）时，转账成功金额随时可能变化
		SuccessNum      int    `json:"success_num"`            // 转账成功的笔数。当批次状态为“PROCESSING”（转账中）时，转账成功笔数随时可能变化
		FailAmount      int    `json:"fail_amount"`            // 转账失败的金额，单位为“分”
		FailNum         int    `json:"fail_num"`               // 转账失败的笔数
		TransferSceneId string `json:"transfer_scene_id"`      // 指定的转账场景ID
	} `json:"transfer_batch"` // 转账批次单基本信息
	TransferDetailList []struct {
		DetailId     string `json:"detail_id"`     // 微信支付系统内部区分转账批次单下不同转账明细单的唯一标识
		OutDetailNo  string `json:"out_detail_no"` // 商户系统内部区分转账批次单下不同转账明细单的唯一标识
		DetailStatus string `json:"detail_status"` // INIT: 初始态。 系统转账校验中 WAIT_PAY: 待确认。待商户确认, 符合免密条件时, 系统会自动扭转为转账中 PROCESSING:转账中。正在处理中，转账结果尚未明确 SUCCESS:转账成功 FAIL:转账失败。需要确认失败原因后，再决定是否重新发起对该笔明细单的转账（并非整个转账批次单）
	} `json:"transfer_detail_list,omitempty"` // 当批次状态为“FINISHED”（已完成），且成功查询到转账明细单时返回。包括微信明细单号、明细状态信息
}

type TransferBatchesOutBatchNoOutBatchNoResult struct {
	Result TransferBatchesOutBatchNoOutBatchNoResponse // 结果
	Body   []byte                                      // 内容
	Http   gorequest.Response                          // 请求
	Err    error                                       // 错误
}

func newTransferBatchesOutBatchNoOutBatchNoResult(result TransferBatchesOutBatchNoOutBatchNoResponse, body []byte, http gorequest.Response, err error) *TransferBatchesOutBatchNoOutBatchNoResult {
	return &TransferBatchesOutBatchNoOutBatchNoResult{Result: result, Body: body, Http: http, Err: err}
}

// TransferBatchesOutBatchNoOutBatchNo 通过商家批次单号查询批次单
// https://pay.weixin.qq.com/docs/merchant/apis/batch-transfer-to-balance/transfer-batch/get-transfer-batch-by-out-no.html
func (c *Client) TransferBatchesOutBatchNoOutBatchNo(ctx context.Context, outBatchNo string, notMustParams ...gorequest.Params) *TransferBatchesOutBatchNoOutBatchNoResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/v3/transfer/batches/out-batch-no/"+outBatchNo, params, http.MethodGet, false)
	if err != nil {
		return newTransferBatchesOutBatchNoOutBatchNoResult(TransferBatchesOutBatchNoOutBatchNoResponse{}, request.ResponseBody, request, err)
	}
	// 定义
	var response TransferBatchesOutBatchNoOutBatchNoResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newTransferBatchesOutBatchNoOutBatchNoResult(response, request.ResponseBody, request, err)
}
