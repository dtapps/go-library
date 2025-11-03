package wechatpayopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type BillFundFlowBillGetResponse struct {
	DownloadUrl string `json:"download_url"` // 哈希类型
	HashType    string `json:"hash_type"`    // 哈希值
	HashValue   string `json:"hash_value"`   // 账单下载地址
}

// BillFundFlowBillGet 申请资金账单API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_4_7.shtml
func (c *Client) BillFundFlowBillGet(ctx context.Context, notMustParams ...*gorequest.Params) (response BillFundFlowBillGetResponse, apiError ApiError, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "v3/bill/fundflowbill", params, http.MethodGet, &response, &apiError)
	return
}
