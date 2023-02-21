package wechatpayapiv3

import (
	"context"
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
	"time"
)

type TransferBatchesOutBatchNoOutBatchNoDetailsOutDetailNoOutDetailNoResponse struct {
	Mchid          string    `json:"mchid"`           // 微信支付分配的商户号
	OutBatchNo     string    `json:"out_batch_no"`    // 商户系统内部的商家批次单号，在商户系统内部唯一
	BatchId        string    `json:"batch_id"`        // 微信批次单号，微信商家转账系统返回的唯一标识
	Appid          string    `json:"appid"`           // 申请商户号的appid或商户号绑定的appid（企业号corpid即为此appid）
	OutDetailNo    string    `json:"out_detail_no"`   // 商户系统内部区分转账批次单下不同转账明细单的唯一标识
	DetailId       string    `json:"detail_id"`       // 微信支付系统内部区分转账批次单下不同转账明细单的唯一标识
	DetailStatus   string    `json:"detail_status"`   // INIT: 初始态。 系统转账校验中 WAIT_PAY: 待确认。待商户确认, 符合免密条件时, 系统会自动扭转为转账中 PROCESSING:转账中。正在处理中，转账结果尚未明确 SUCCESS:转账成功 FAIL:转账失败。需要确认失败原因后，再决定是否重新发起对该笔明细单的转账（并非整个转账批次单）
	TransferAmount int       `json:"transfer_amount"` // 转账金额单位为“分”
	TransferRemark string    `json:"transfer_remark"` // 单条转账备注（微信用户会收到该备注），UTF8编码，最多允许32个字符
	FailReason     string    `json:"fail_reason"`     // 如果转账失败则有失败原因
	Openid         string    `json:"openid"`          // 商户appid下，某用户的openid
	UserName       string    `json:"user_name"`       // 收款方姓名。采用标准RSA算法，公钥由微信侧提供 商户转账时传入了收款用户姓名、查询时会返回收款用户姓名
	InitiateTime   time.Time `json:"initiate_time"`   // 转账发起的时间，按照使用rfc3339所定义的格式，格式为YYYY-MM-DDThh:mm:ss+TIMEZONE
	UpdateTime     time.Time `json:"update_time"`     // 明细最后一次状态变更的时间，按照使用rfc3339所定义的格式，格式为YYYY-MM-DDThh:mm:ss+TIMEZONE
}

type TransferBatchesOutBatchNoOutBatchNoDetailsOutDetailNoOutDetailNoResult struct {
	Result TransferBatchesOutBatchNoOutBatchNoDetailsOutDetailNoOutDetailNoResponse // 结果
	Body   []byte                                                                   // 内容
	Http   gorequest.Response                                                       // 请求
	Err    error                                                                    // 错误
}

func newTransferBatchesOutBatchNoOutBatchNoDetailsOutDetailNoOutDetailNoResult(result TransferBatchesOutBatchNoOutBatchNoDetailsOutDetailNoOutDetailNoResponse, body []byte, http gorequest.Response, err error) *TransferBatchesOutBatchNoOutBatchNoDetailsOutDetailNoOutDetailNoResult {
	return &TransferBatchesOutBatchNoOutBatchNoDetailsOutDetailNoOutDetailNoResult{Result: result, Body: body, Http: http, Err: err}
}

// TransferBatchesOutBatchNoOutBatchNoDetailsOutDetailNoOutDetailNo 通过商家明细单号查询明细单
// https://pay.weixin.qq.com/docs/merchant/apis/batch-transfer-to-balance/transfer-detail/get-transfer-detail-by-out-no.html
func (c *Client) TransferBatchesOutBatchNoOutBatchNoDetailsOutDetailNoOutDetailNo(ctx context.Context, outBatchNo, outDetailNo string, notMustParams ...gorequest.Params) *TransferBatchesOutBatchNoOutBatchNoDetailsOutDetailNoOutDetailNoResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/v3/transfer/batches/out-batch-no/"+outBatchNo+"/details/out-detail-no/"+outDetailNo, params, http.MethodGet, false)
	if err != nil {
		return newTransferBatchesOutBatchNoOutBatchNoDetailsOutDetailNoOutDetailNoResult(TransferBatchesOutBatchNoOutBatchNoDetailsOutDetailNoOutDetailNoResponse{}, request.ResponseBody, request, err)
	}
	// 定义
	var response TransferBatchesOutBatchNoOutBatchNoDetailsOutDetailNoOutDetailNoResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newTransferBatchesOutBatchNoOutBatchNoDetailsOutDetailNoOutDetailNoResult(response, request.ResponseBody, request, err)
}
