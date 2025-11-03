package wechatpayopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type billSubMerchantFundFlowBillGetResponse struct {
	DownloadBillCount int `json:"download_bill_count"` // 下载信息总数
	DownloadBillList  []struct {
		BillSequence int    `json:"bill_sequence"` // 账单文件序号
		DownloadUrl  string `json:"download_url"`  // 下载地址
		EncryptKey   string `json:"encrypt_key"`   // 加密密钥
		HashType     string `json:"hash_type"`     // 哈希类型
		HashValue    string `json:"hash_value"`    // 哈希值
		Nonce        string `json:"nonce"`         // 随机字符串
	} `json:"download_bill_list"` // 下载信息明细
}

// billSubMerchantFundFlowBillGet 申请单个子商户资金账单API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_4_12.shtml
func (c *Client) billSubMerchantFundFlowBillGet(ctx context.Context, notMustParams ...*gorequest.Params) (response billSubMerchantFundFlowBillGetResponse, apiError ApiError, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("sub_mchid", c.GetSubMchId()) // 子商户号

	// 请求
	err = c.request(ctx, "v3/bill/sub-merchant-fundflowbill", params, http.MethodGet, &response, &apiError)
	return
}
