package wechatpayopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type BillTradeBillGetResponse struct {
	DownloadUrl string `json:"download_url"` // 哈希类型
	HashType    string `json:"hash_type"`    // 哈希值
	HashValue   string `json:"hash_value"`   // 账单下载地址
}

// BillTradeBillGet 申请交易账单API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_4_6.shtml
func (c *Client) BillTradeBillGet(ctx context.Context, notMustParams ...*gorequest.Params) (response BillTradeBillGetResponse, apiError ApiError, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("sub_mchid", c.GetSubMchId()) // 子商户号

	// 请求
	err = c.request(ctx, "/v3/bill/tradebill", params, http.MethodGet, &response, &apiError)
	return
}
