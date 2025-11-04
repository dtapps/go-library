package wechatpayapiv3

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type TransferDetailElectronicReceiptsPostResponse struct {
	AcceptType      string `json:"accept_type"`      // 电子回单受理类型：BATCH_TRANSFER：批量转账明细电子回单 TRANSFER_TO_POCKET：企业付款至零钱电子回单 TRANSFER_TO_BANK：企业付款至银行卡电子回单
	OutBatchNo      string `json:"out_batch_no"`     // 需要电子回单的批量转账明细单所在的转账批次的单号，该单号为商户申请转账时生成的商户单号。受理类型为BATCH_TRANSFER时该单号必填，否则该单号留空。
	OutDetailNo     string `json:"out_detail_no"`    // 该单号为商户申请转账时生成的商家转账明细单号。 1.受理类型为BATCH_TRANSFER时填写商家批量转账明细单号。2. 受理类型为TRANSFER_TO_POCKET或TRANSFER_TO_BANK时填写商家转账单号。
	SignatureNo     string `json:"signature_no"`     // 电子回单受理单号，受理单据的唯一标识
	SignatureStatus string `json:"signature_status"` // 枚举值： ACCEPTED:已受理，电子签章已受理成功 FINISHED:已完成。电子签章已处理完成
	HashType        string `json:"hash_type"`        // 电子回单文件的hash方法，回单状态为：FINISHED时返回
	HashValue       string `json:"hash_value"`       // 电子回单文件的hash值，用于下载之后验证文件的完整、正确性，回单状态为：FINISHED时返回
	DownloadUrl     string `json:"download_url"`     // 电子回单文件的下载地址，回单状态为：FINISHED时返回。URL有效时长为10分钟，10分钟后需要重新去获取下载地址（但不需要走受理）
}

// TransferDetailElectronicReceiptsPost 受理转账明细电子回单API
// https://pay.weixin.qq.com/docs/merchant/apis/batch-transfer-to-balance/electronic-receipt-api/create-electronic-receipt.html
func (c *Client) TransferDetailElectronicReceiptsPost(ctx context.Context, notMustParams ...*gorequest.Params) (response TransferDetailElectronicReceiptsPostResponse, apiError ApiError, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.DoRequest(ctx, "v3/transfer-detail/electronic-receipts", params, http.MethodPost, false, &response)
	return
}
